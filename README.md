### README.md


# Projeto Eicon-Teste

Este projeto é um exemplo de aplicação Go utilizando vários pacotes populares como Gin, Gorm e Viper. Siga as instruções abaixo para configurar e executar o projeto em ambientes Linux e Windows.

## Pré-requisitos

- [Go](https://golang.org/dl/)
- Git Bash (para usuários do Windows)
- GCC e libc6-dev (para usuários de Linux)

## Instalação do Go

### Linux

1. **Atualize e faça upgrade no sistema:**
   ```bash
   sudo apt update
   sudo apt upgrade
   ```

2. **Baixe o tarball do Go:**
   ```bash
   wget https://go.dev/dl/go1.20.5.linux-amd64.tar.gz
   ```

3. **Extraia o tarball:**
   ```bash
   sudo tar -C /usr/local -xzf go1.20.5.linux-amd64.tar.gz
   ```

4. **Configure o ambiente Go:**
   Adicione as seguintes linhas ao arquivo `~/.profile`:
   ```bash
   echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.profile
   source ~/.profile
   ```

### Windows

1. **Baixe o instalador do Go:**
   Acesse [Go Downloads](https://golang.org/dl/) e baixe o instalador para Windows.

2. **Execute o instalador e siga as instruções.**

3. **Configure o ambiente Go no Git Bash:**
   Adicione as seguintes linhas ao arquivo `~/.bash_profile`:
   ```bash
   echo "export PATH=$PATH:/c/Go/bin" >> ~/.bash_profile
   source ~/.bash_profile
   ```

## Configuração do Ambiente

Para garantir que o ambiente esteja configurado corretamente, execute o script `setup_env.sh`. Este script configura as variáveis de ambiente e instala as dependências necessárias.

### Passos para Configuração

1. **Clone o repositório:**

   ```bash
   git clone <URL_DO_REPOSITORIO>
   cd eicon-teste
   ```

2. **Torne o script executável (apenas para Linux):**

   ```bash
   chmod +x setup_env.sh
   ```

3. **Execute o script de configuração:**

   ```bash
   ./setup_env.sh
   ```

### Explicação do Script

O script `setup_env.sh` realiza as seguintes tarefas:

- Para Linux:
  - Atualiza a lista de pacotes.
  - Instala `gcc` e `libc6-dev`.
  - Define a variável de ambiente `CGO_ENABLED=1`.
  - Adiciona `CGO_ENABLED=1` ao arquivo `~/.bashrc`.

- Para Windows (usando Git Bash):
  - Define a variável de ambiente `CGO_ENABLED=1`.
  - Adiciona `CGO_ENABLED=1` ao arquivo `~/.bash_profile`.

- Limpa o cache do Go.
- Executa a aplicação Go (`main.go`).

## Executando a Aplicação

Após a configuração do ambiente, você pode executar a aplicação com o comando:

```bash
go run main.go
```

### Solução de Problemas

Se você encontrar o erro `Binary was compiled with 'CGO_ENABLED=0', go-sqlite3 requires cgo to work`, certifique-se de que a variável de ambiente `CGO_ENABLED` está definida como `1` e que todas as dependências estão instaladas corretamente.

## Nota sobre o Projeto

Este projeto poderia ter sido feito em Java, mas optei por usar Go para demonstrar minhas habilidades diversificadas em várias linguagens e tecnologias. Usar Go mostra uma abordagem eficiente e moderna para desenvolvimento de aplicações web.

## Contato

Para dúvidas ou suporte, entre em contato com Micael Santana em micaelparadox@gmail.com.
