#!/usr/bin/env python3
"""
RustDesk Customization Script
Replaces the closed-source 'generator' binary.
Applies customizations to the RustDesk source tree before build.

Usage:
  python3 customize.py <custom_config_json> [--os windows] [--arch x86_64]

Environment variables set by this script (written to $GITHUB_ENV if present):
  APP_NAME, CUSTOM_PASSWORD, CUSTOM_SALT, CONN_TYPE
"""

import json
import os
import sys
import re
import shutil
from pathlib import Path


def find_file(root, name):
    for path in Path(root).rglob(name):
        return str(path)
    return None


def replace_in_file(filepath, old, new):
    if not os.path.exists(filepath):
        print(f"  [WARN] File not found: {filepath}")
        return False
    with open(filepath, 'r', encoding='utf-8', errors='replace') as f:
        content = f.read()
    if old not in content:
        print(f"  [WARN] Pattern not found in {filepath}: {repr(old)[:80]}")
        return False
    content = content.replace(old, new)
    with open(filepath, 'w', encoding='utf-8') as f:
        f.write(content)
    print(f"  [OK] Replaced in {filepath}")
    return True


def set_env(name, value):
    """Set environment variable for GitHub Actions."""
    if value is not None:
        github_env = os.environ.get('GITHUB_ENV', '')
        if github_env:
            with open(github_env, 'a') as f:
                f.write(f"{name}={value}\n")
        os.environ[name] = str(value)
        print(f"  [ENV] {name}={value}")


def apply_config(config_path, config):
    """Apply configuration to hbb_common/src/config.rs"""
    app_name = config.get('app-name', 'RustDesk')
    password = config.get('password', '')
    salt = config.get('salt', '')
    conn_type = config.get('conn-type', 'both')

    if not os.path.exists(config_path):
        print(f"  [WARN] Config file not found: {config_path}")
        return

    with open(config_path, 'r', encoding='utf-8', errors='replace') as f:
        content = f.read()

    # 1. Override APP_NAME
    app_name_escaped = app_name.replace('\\', '\\\\').replace('"', '\\"')
    content = re.sub(
        r'pub const APP_NAME:.*=.*".*";',
        f'pub const APP_NAME: &str = "{app_name_escaped}";',
        content,
    )
    content = re.sub(
        r'pub const APP_NAME_READABLE:.*=.*".*";',
        f'pub const APP_NAME_READABLE: &str = "{app_name_escaped}";',
        content,
    )

    # 2. Override default password/salt in config
    if password:
        content = re.sub(
            r'pub const PASSWORD:.*=.*".*";',
            f'pub const PASSWORD: &str = "{password}";',
            content,
        )

    if salt:
        content = re.sub(
            r'pub const SALT:.*=.*".*";',
            f'pub const SALT: &str = "{salt}";',
            content,
        )

    # 3. Set connection type
    if conn_type in ('incoming', 'outgoing', 'both'):
        content = re.sub(
            r'pub const CONN_TYPE:.*=.*".*";',
            f'pub const CONN_TYPE: &str = "{conn_type}";',
            content,
        )

    with open(config_path, 'w', encoding='utf-8') as f:
        f.write(content)
    print(f"  [CONFIG] Applied to {config_path}")


def apply_feature_flags(features_path, config):
    """Apply feature toggles to RustDesk features config."""
    if not os.path.exists(features_path):
        print(f"  [WARN] Features file not found: {features_path}")
        return

    flags = {
        'disable-installation': ('disable_installation', 'N'),
        'disable-settings': ('disable_settings', 'N'),
        'disable-account': ('disable_account', 'N'),
        'disable-ab': ('disable_ab', 'N'),
        'disable-tcp-listen': ('disable_tcp_listen', 'N'),
    }

    for key, (feature_name, _) in flags.items():
        value = config.get(key, 'N')
        set_env(feature_name.upper(), 'Y' if value == 'Y' else 'N')


def main():
    if len(sys.argv) < 2:
        print("Usage: python3 customize.py <custom_config_json> [--os windows] [--arch x86_64]")
        sys.exit(1)

    custom_json = sys.argv[1]
    try:
        config = json.loads(custom_json)
    except json.JSONDecodeError:
        print("Error: invalid JSON")
        sys.exit(1)

    os_platform = next((sys.argv[i+1] for i, a in enumerate(sys.argv) if a == '--os'), 'windows')
    arch = next((sys.argv[i+1] for i, a in enumerate(sys.argv) if a == '--arch'), 'x86_64')

    print(f"[Customize] OS={os_platform}, ARCH={arch}")
    print(f"[Customize] Config: {json.dumps(config, indent=2)}")

    app_name = config.get('app-name', 'RustDesk')
    set_env('APP_NAME', app_name)

    # Apply to config.rs
    config_path = find_file('.', 'config.rs')
    if config_path and 'hbb_common' in config_path:
        apply_config(config_path, config)
    else:
        # Try specific path
        specific = 'libs/hbb_common/src/config.rs'
        if os.path.exists(specific):
            apply_config(specific, config)
        else:
            print("  [WARN] Could not find config.rs")

    # Set feature flags
    set_env('DISABLE_INSTALLATION', config.get('disable-installation', 'N'))
    set_env('DISABLE_SETTINGS', config.get('disable-settings', 'N'))
    set_env('DISABLE_ACCOUNT', config.get('disable-account', 'N'))
    set_env('DISABLE_AB', config.get('disable-ab', 'N'))
    set_env('DISABLE_TCP_LISTEN', config.get('disable-tcp-listen', 'N'))

    # Save override settings
    override = config.get('override-settings', {})
    if override:
        override_path = 'override_settings.json'
        with open(override_path, 'w') as f:
            json.dump(override, f)
        set_env('OVERRIDE_SETTINGS_FILE', override_path)
        print(f"  [OVERRIDE] Settings saved to {override_path}")

    # Default settings
    default = config.get('default-settings', {})
    if default:
        default_path = 'default_settings.json'
        with open(default_path, 'w') as f:
            json.dump(default, f)
        set_env('DEFAULT_SETTINGS_FILE', default_path)
        print(f"  [DEFAULT] Settings saved to {default_path}")

    # Rename executable references if app_name differs from RustDesk
    if app_name != 'RustDesk':
        # Modify Cargo.toml if needed
        cargo_toml = find_file('.', 'Cargo.toml')
        if cargo_toml:
            with open(cargo_toml, 'r', encoding='utf-8', errors='replace') as f:
                content = f.read()
            # Add/update the name in the root package
            content = re.sub(
                r'^name = "rustdesk"',
                f'name = "{app_name.lower().replace(" ", "-")}"',
                content,
                flags=re.MULTILINE,
            )
            with open(cargo_toml, 'w', encoding='utf-8') as f:
                f.write(content)
            print(f"  [CARGO] Updated package name in Cargo.toml")

        # Modify build.py for the output
        build_py = 'build.py'
        if os.path.exists(build_py):
            with open(build_py, 'r', encoding='utf-8', errors='replace') as f:
                content = f.read()
            content = content.replace('rustdesk.exe', f'{app_name}.exe')
            content = content.replace('rustdesk-portable-packer.exe', f'{app_name}-portable-packer.exe')
            with open(build_py, 'w', encoding='utf-8') as f:
                f.write(content)
            print(f"  [BUILD] Updated executable names in build.py")

    print("[Customize] Done!")


if __name__ == '__main__':
    main()
