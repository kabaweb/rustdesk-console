param(
    [string]$AppName = "bmdesk",
    [string]$Arch = "x86_64",
    [string]$ConnType = "",
    [string]$CustomConfig = "{}",
    [string]$SourcePath = "D:\Python\bmdesk",
    [string]$OutputDir = ".\output"
)

$ErrorActionPreference = "Stop"
Push-Location $SourcePath

try {
    Write-Host "=== [1/4] Set env vars ==="
    $env:APP_NAME = $AppName
    python -c "import json; print(json.dumps(json.loads('$CustomConfig'), indent=2))"

    Write-Host "=== [2/4] Apply customizations ==="
    if ($CustomConfig -ne "{}") {
        python "$PSScriptRoot\..\customize\apply.py" '$CustomConfig'
    }

    Write-Host "=== [3/4] Build Rust + Flutter ==="
    $buildArgs = "--portable --flutter"
    if ($ConnType -eq "incoming" -or $ConnType -eq "outgoing") {
        $buildArgs += " --conn-type $ConnType"
    }
    python build.py @($buildArgs -split ' ')

    Write-Host "=== [4/4] Collect output ==="
    $outDir = Join-Path (Resolve-Path $OutputDir) $AppName
    New-Item -ItemType Directory -Force -Path $outDir | Out-Null

    # Portable installer
    Get-ChildItem . -Filter "${AppName}*.exe" | Copy-Item -Destination $outDir -Force
    if (Test-Path "target\release\rustdesk-portable-packer.exe") {
        Copy-Item "target\release\rustdesk-portable-packer.exe" "$outDir\" -Force
    }

    # Full Release tree as fallback
    $releaseTree = "flutter\build\windows\x64\runner\Release"
    if (Test-Path $releaseTree) {
        Copy-Item -Recurse "$releaseTree\*" "$outDir\" -Force
    }

    Write-Host ""
    Write-Host "=== BUILD COMPLETE ==="
    Write-Host "Output: $outDir"
    Get-ChildItem $outDir | ForEach-Object { Write-Host "  $_" }
} finally {
    Pop-Location
}
