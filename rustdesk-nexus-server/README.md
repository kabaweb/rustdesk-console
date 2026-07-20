# RustDesk Nexus Server

Servidor self-hosted que substitui o `api.databk.top`, permitindo gerar clientes RustDesk personalizados
a partir do **seu próprio fork** do RustDesk.

## Arquitetura

```
rustdesk-console (NestJS)         ← seu painel admin
        │
        │ NEXUS_BASE_URL
        ▼
rustdesk-nexus-server (Go)        ← este projeto
        │
        │ GitHub API (workflow_dispatch)
        ▼
seu-fork/rustdesk-generator       ← GitHub Actions
        │
        │ checkout + customize
        ▼
seu-fork/rustdesk                 ← seu código fonte personalizado
        │
        │ build (flutter + rust)
        ▼
artifacts (.exe, .msi, etc.)     ← disponível para download
```

## Pré-requisitos

| Requisito | Descrição |
|-----------|-----------|
| Go 1.22+ | Para compilar o servidor |
| GitHub Token | [Criar token](https://github.com/settings/tokens) com scopes: `repo`, `workflow` |
| Fork do rustdesk-generator | [databk/rustdesk-generator](https://github.com/databk/rustdesk-generator) |
| Fork do RustDesk | Seu fork com as modificações desejadas (base: tag `1.4.9`) |

---

## Passo a passo

### 1. Criar token do GitHub

1. Acesse https://github.com/settings/tokens
2. Generate new token (classic)
3. Escopo: `repo` (completo) + `workflow`
4. Copie o token gerado

### 2. Preparar o fork do rustdesk-generator

```bash
# Fork o repo no GitHub (botão Fork em github.com/databk/rustdesk-generator)
# Depois clone seu fork:
git clone https://github.com/SEU_USER/rustdesk-generator.git
cd rustdesk-generator

# Copie os arquivos deste projeto para dentro do fork:
cp ../rustdesk-console/rustdesk-nexus-server/workflows/build-custom-windows.yml .github/workflows/
cp ../rustdesk-console/rustdesk-nexus-server/customize/apply.py customize/

# Commit e push
git add .github/workflows/build-custom-windows.yml customize/apply.py
git commit -m "Add custom build workflow and customization script"
git push
```

### 3. Configurar o servidor Nexus

```bash
cd rustdesk-console/rustdesk-nexus-server
cp .env.example .env
```

Edite o arquivo `.env`:

```env
# Porta do servidor
NEXUS_PORT=8090

# JWT secret - gere uma string aleatória
NEXUS_JWT_SECRET=SUA_STRING_ALEATORIA_AQUI

# Token do GitHub (passo 1)
GITHUB_TOKEN=ghp_xxxxxxxxxxxxxxxxxxxx

# Seu fork do rustdesk-generator (passo 2)
GITHUB_OWNER=SEU_USER
GITHUB_REPO=rustdesk-generator
GITHUB_WORKFLOW_FILE=build-custom-windows.yml

# SEU fork do RustDesk (fonte do build)
RUSTDESK_REPO=SEU_USER/rustdesk
RUSTDESK_REF=refs/tags/1.4.9

# URL pública (se acessível externamente, ex: https://nexus.seudominio.com)
NEXUS_PUBLIC_URL=

# Limite de builds por mês (0 = ilimitado)
NEXUS_MAX_BUILDS_PER_MONTH=15
```

### 4. Iniciar o servidor

**Opção A - Go (recomendado para Windows):**

```bash
# Instalar Go: https://go.dev/dl/
go mod tidy
go run .
```

**Opção B - Docker:**

```bash
docker-compose up -d
```

O servidor inicia em `http://localhost:8090`.

### 5. Conectar o console ao servidor

No arquivo `.env` do seu `rustdesk-console`, adicione:

```env
NEXUS_BASE_URL=http://localhost:8090
```

Reinicie o console:

```bash
npm run start:dev
```

### 6. Autorizar o console no servidor

1. No painel do console, vá até a seção de geração de cliente (Nexus)
2. Clique em "Vincular conta" — o console vai chamar o servidor Nexus
3. Abra a URL de autorização retornada no navegador
4. Clique em "Authorize" na página simples que aparece
5. O console detecta a autorização automaticamente

### 7. Gerar um cliente personalizado

No painel do console, preencha:

| Campo | Descrição | Exemplo |
|-------|-----------|---------|
| **OS** | Sistema operacional | `windows` |
| **Arch** | Arquitetura | `x86_64` (ou `aarch64`, `x86`) |
| **app-name** | Nome do executável | `MeuApp` |
| **password** | Senha padrão de conexão | `minhasenha` |
| **salt** | Salt da senha | `meusal` |
| **conn-type** | Direção da conexão | `both`, `incoming`, `outgoing` |
| **disable-installation** | Desabilitar instalação | `Y` ou `N` |
| **disable-settings** | Desabilitar configurações | `Y` ou `N` |
| **disable-account** | Desabilitar conta | `Y` ou `N` |
| **disable-ab** | Desabilitar catálogo | `Y` ou `N` |
| **disable-tcp-listen** | Desabilitar TCP listen | `Y` ou `N` |
| **override-settings** | Configurações sobrescritas (JSON) | `{"key": "value"}` |
| **default-settings** | Configurações padrão (JSON) | `{"key": "value"}` |

### 8. Baixar o cliente

O console faz polling a cada 10 segundos. Quando o build terminar (15-30 min no GitHub Actions),
os arquivos aparecem para download no painel.

---

## Como funciona o script de customização

O arquivo `customize/apply.py` substitui o binário `generator.exe` (fechado) do projeto original.
Ele faz as seguintes modificações no código fonte do RustDesk **antes** da compilação:

| Customização | Arquivo modificado | O que faz |
|---|---|---|
| `app-name` | `libs/hbb_common/src/config.rs` | Altera `APP_NAME` e `APP_NAME_READABLE` |
| `app-name` | `build.py` | Renomeia executáveis de saída |
| `app-name` | `Cargo.toml` | Altera o nome do pacote |
| `password` | `libs/hbb_common/src/config.rs` | Define senha padrão |
| `salt` | `libs/hbb_common/src/config.rs` | Define salt padrão |
| `conn-type` | `libs/hbb_common/src/config.rs` | Define direção de conexão |
| `disable-*` | Variáveis de ambiente | Setadas para o build |
| `override-settings` | `override_settings.json` | Salvo como arquivo JSON |
| `default-settings` | `default_settings.json` | Salvo como arquivo JSON |

---

## Troubleshooting

### O build falha no GitHub Actions

1. Verifique os logs em: `https://github.com/SEU_USER/rustdesk-generator/actions`
2. Causas comuns:
   - Tag `1.4.9` não existe no seu fork → ajuste `RUSTDESK_REF`
   - Falta de dependências no runner → compare com o workflow original
   - Erro no script de customização → veja logs do step "Apply customizations"

### O servidor não dispara o workflow

1. Verifique se `GITHUB_TOKEN` tem permissão `workflow`
2. Verifique se o arquivo YAML está em `.github/workflows/` no fork
3. O nome do arquivo no `.env` (`GITHUB_WORKFLOW_FILE`) deve bater exatamente

### O console não conecta ao servidor

1. Teste: `curl http://localhost:8090/health`
2. Verifique se `NEXUS_BASE_URL` no console aponta para `http://localhost:8090`
3. Se o servidor está em Docker e o console em host, use `host.docker.internal`

### Erro "missing token" ao gerar build

O console precisa estar vinculado. No painel:
1. Vá em Nexus → Vincular conta
2. Abra a URL retornada no navegador
3. Após autorizar, o console detecta automaticamente

---

## Personalização avançada

### Adicionar mais sistemas operacionais

1. Crie workflows adicionais em `.github/workflows/` (ex: `build-custom-linux.yml`)
2. Use os workflows originais do `rustdesk-generator` como base (substitua `generator.exe` pelo `apply.py`)
3. No console, modifique `NexusGenerateDto` em `src/modules/nexus/dto/nexus-client.dto.ts` para aceitar `'linux'`, `'macos'`, etc.

### Usar sem GitHub Actions (build local)

Se preferir build local em vez de GitHub Actions:
1. Remova as configurações `GITHUB_*` do `.env`
2. O servidor registra o build como `pending` mas não dispara workflow
3. Implemente um worker que execute o build localmente (Docker, VM, etc.)
4. Atualize o status via API interna

### Autenticação real com GitHub OAuth

O auth atual é simplificado (qualquer um se autentica). Para OAuth real:
1. Crie um GitHub OAuth App em https://github.com/settings/developers
2. Modifique `api/auth.go` para usar a API OAuth do GitHub
3. Adicione `GITHUB_CLIENT_ID` e `GITHUB_CLIENT_SECRET` ao `.env`

---

## API Endpoints

O servidor implementa o mesmo contrato do `api.databk.top`:

| Método | Rota | Auth | Descrição |
|--------|------|------|-----------|
| `GET` | `/v1/auth/github/login` | Não | Cria sessão de login |
| `GET` | `/v1/auth/github/status` | Não | Verifica status do login |
| `GET` | `/v1/auth/github/callback` | Não | Callback OAuth (página HTML) |
| `POST` | `/v1/client/generate` | JWT | Submete build |
| `GET` | `/v1/client/generate/{uuid}` | JWT | Status do build |
| `GET` | `/v1/client/download/{uuid}/{file}` | JWT | Download do artefato |
| `GET` | `/health` | Não | Health check |
