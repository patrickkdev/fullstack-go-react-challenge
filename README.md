# 📋 Projeto Recrutamento & Seleção

Este repositório contém a implementação de um sistema de **Recrutamento e Seleção**, com **backend** em Go (framework Gin) e **frontend** em TypeScript (React + Vite). O objetivo é gerir candidatos, vagas e etapas de processo seletivo de forma simples e organizada.

1. Tela de cadastro (EMAIL e SENHA)
2. Tela de login (EMAIL e SENHA)
3. Tela quando logado (Minhas vagas / Minhas candidaturas)
4. Registro de vagas
5. Cadastro em vaga
6. Busca de vagas

Atualizar a tela não posso perder o login, quando logado não posso entrar na tela de login e registro.

## NÃO UTILIZE NEXT.JS


---

## 📚 Tecnologias

### Backend
- [Go](https://golang.org/)  
- [Gin Web Framework](https://github.com/gin-gonic/gin)  
- [GORM](https://gorm.io/)  
- PostgreSQL  

### Frontend
- [TypeScript](https://www.typescriptlang.org/)  
- [React](https://reactjs.org/)  
- [Vite](https://vitejs.dev/)  
- [Axios](https://github.com/axios/axios)  

---

## 🐳 Docker Setup

Este projeto inclui uma configuração Docker Compose para facilitar o desenvolvimento local com PostgreSQL.

### Pré-requisitos
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/)

### Configuração
1. Copie o arquivo de exemplo de variáveis de ambiente:
   ```bash
   cp backend/.env.example backend/.env
   ```

2. Execute o Docker Compose:
   ```bash
   docker-compose up --build
   ```

3. A aplicação estará disponível em:
   - Frontend: http://localhost:5173
   - Backend: http://localhost:4000
   - PostgreSQL: localhost:5432

### Comandos Úteis
- `docker-compose up -d` - Executar em background
- `docker-compose down` - Parar e remover containers
- `docker-compose logs` - Ver logs
- `docker-compose exec postgres psql -U postgres -d recruiting` - Acessar o banco de dados

---