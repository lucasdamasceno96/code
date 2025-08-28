# --- Estágio 1: Builder ---
# Usamos uma imagem oficial do Go para construir a aplicação.
# O 'alpine' a torna menor. Nomeamos este estágio como 'builder'.
FROM golang:1.19-alpine AS builder

# É necessário para o driver do SQLite.
RUN apk add --no-cache gcc musl-dev

# Define o diretório de trabalho dentro do contêiner.
WORKDIR /app

# Copia os arquivos de módulo e baixa as dependências.
# Isso aproveita o cache do Docker.
COPY go.mod go.sum ./
RUN go mod download

# Copia todo o código-fonte da aplicação.
COPY . .

# Constrói a aplicação.
# CGO_ENABLED=1 é obrigatório para o driver do SQLite.
# -o /url-shortener especifica que o executável será criado na raiz com esse nome.
RUN CGO_ENABLED=1 GOOS=linux go build -o /url-shortener cmd/api/main.go


# --- Estágio 2: Final ---
# Começamos com uma imagem base mínima. Alpine é pequena e segura.
FROM alpine:latest

# Define o diretório de trabalho.
WORKDIR /app

# Cria um diretório para o banco de dados.
# Não copiamos o banco de dados para a imagem, pois ele deve ser persistido com um volume.
RUN mkdir data

# Copia APENAS o executável compilado do estágio 'builder'.
# Nenhum código-fonte ou ferramenta de build é incluído na imagem final.
COPY --from=builder /url-shortener .

# Expõe a porta que a nossa aplicação usa.
EXPOSE 8080

# O comando para iniciar a aplicação quando o contêiner for executado.
CMD ["./url-shortener"]