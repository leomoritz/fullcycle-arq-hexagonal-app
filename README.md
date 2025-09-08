
# Arquitetura Hexagonal (Ports and Adapters)

## Introdução

### Pontos importantes sobre arquitetura
- Crescimento sustentável
- Software precisa gerar valor ao longo do tempo
- Software deve ser desenhado por você e não pelo seu framework
- Componentes devem ser encaixáveis e substituíveis com facilidade
- Arquitetar software é diferente de escrever software. O software pode morrer não por ser mal escrito, mas por ter sido mal desenhado
- "Arquitetura diz respeito com o futuro do seu software. CRUD qualquer um faz!"

## Ciclo de Vida de Muitos Projetos

### Fase 1
- 🗄️ Banco de dados
- 📝 Cadastros
- ✅ Validações
- 🌐 Servidor web
- 🧭 Controllers
- 🖼️ Views
- 🔐 Autenticação
- 📤 Upload de arquivos

### Fase 2
- 📊 Regras de negócio
- 🔗 Criação de APIs
- 🔄 Consumo de APIs
- 🔒 Autorização
- 📈 Relatórios
- 🧾 Logs

### Fase 3
- 🚀 Mais acessos
- 💾 Upgrades de hardware (escala vertical)
- 🧠 Cache
- 🤝 Consumir APIs de parceiros
- 📊 Regras de parceiros
- 📉 Relatórios

### Fase 4
- 🚀 Mais acessos
- 💾 Upgrade de hardware (escala vertical)
- 🗃️ Banco de dados para relatórios
- 🧾 Comandos
- 🔄 V2 da API

### Fase 5
- ⚖️ Escala horizontal
- 🧑‍💻 Sessões
- 📤 Uploads
- 🛠️ Refatoração
- 📈 Autoscaling
- 🔁 CI/CD

### Fase 6
- 🔍 GraphQL
- 🐞 Bugs constantes
- 📉 Logs? Ops
- 🔗 Integração com outro serviço (Ex: CRM)
- 🔄 Migração para React

### Fase 7
- ❌ Inconsistência na integração com outro serviço (Ex: CRM)
- 📦 Containers
- 🔁 CI/CD para adaptar com container
- 🧠 Memória
- 📉 Logs
- 🧹 Se livrar do legado

### Fase 8
- 🧩 Microserviços
- 🗃️ DB compartilhado
- 🔍 Problemas com tracing
- 🐢 Lentidão
- 💸 Custo elevado

### Fase 9
- ☸️ Kubernetes
- 🔁 CI/CD para adaptar ao Kubernetes
- 📬 Mensageria
- ❗ Perda de mensagens
- 🧑‍🏫 Consultorias para ajudar

### Fase 10
- 🧠 Use a imaginação!

## Principais Problemas
- 🔮 Falta de visão de futuro
- 🧱 Limites mal definidos
- 🔄 Dificuldade na troca e adição de componentes
- ⚖️ Escala horizontal
- ⚙️ Otimizações frequentes
- 🔄 Preparado para mudanças

## Reflexões
> Antes de seguir com decisões arquiteturais, vale refletir sobre os impactos reais no time e no cliente:

- Está sendo doloroso para o developer?
- Poderia ter sido evitado?
- O software está se pagando?
- A relação com o cliente está boa?
- O cliente terá prejuízo com a brusca mudança arquitetural?
- Em qual momento tudo se perdeu? A resposta é: No dia 1 e nos dias subsequentes, um dia de cada vez.
- Se você fosse novo na equipe, como avaliaria as decisões tomadas até aqui?

## Arquitetura vs Design

> "Atividades relacionadas à arquitetura de software são sempre de design. Entretanto, nem todas atividades de design são sobre arquitetura. O objetivo primário da arquitetura de software é garantir que os atributos de qualidade, restrições de alto nível e os objetivos do negócio, sejam atendidos pelo sistema. Qualquer decisão de design que não tenha relação com este objetivo não é arquitetural. Todas as decisões de design para um componente que não sejam 'visíveis' fora dele, geralmente, também não são."
> — Elemar Jr. ([Fonte](https://www.eximiaco.tech/pt/2020/01/08/quais-sao-as-diferencas-entre-arquitetura-e-desing-de-software/))

- Arquitetura é a estratégia definida para garantir que o software será entregue com qualidade.
- O design são os detalhes de como as coisas serão feitas para atender com qualidade o que foi definido pela arquitetura.

## Arquitetura Hexagonal (Ports and Adapters)

> "Allow an application to equally be driven by users, programs, automated test or batch scripts, and to be developed and tested in isolation from its eventual run-time devices and databases." — Cockburn

> "O termo 'Arquitetura Hexagonal' está muito mais ligado com decisões de design de software do que necessariamente de arquitetura." — Wesley Williams

- Na arquitetura hexagonal, a aplicação (core) fica no centro (hexágono) e qualquer acesso externo é feito através de adaptadores criados.
- O hexágono é uma metáfora visual que representa os pontos de conexão entre o núcleo da aplicação e seus adaptadores externos.
- O cliente fica no lado esquerdo (REST, CLI, GRPC, GraphQL, UI, etc) e o servidor no lado direito (DB, Redis, Filesystem, Lambda, API externa, etc).

### Características
- Definição de limites e proteção nas regras da aplicação
- Componentização e desacoplamento:
  - Logs
  - Cache
  - Upload
  - Banco de dados
  - Comandos
  - Filas
  - HTTP / APIs / GraphQL
- Facilidade na quebra para microsserviços

### Princípio Fundamental
- O principal conceito da arquitetura hexagonal é o Princípio da Inversão de Dependência (DIP - Dependency Inversion Principle):
  - Módulos de alto nível não devem depender de módulos de baixo nível. Ambos devem depender de abstrações.
  - Abstrações não devem depender de detalhes. Detalhes devem depender de abstrações.

- Não há padrão estabelecido de como o código deve ser organizado, mas quanto mais desacoplado for o seu código, melhor.

## Especificação Técnica do Projeto
- Go 1.16 (Entenda mais sobre *ponteiros na Golang* [aqui](https://www.youtube.com/watch?v=-FiBp1OeZF0))
- SQLite 3
- Docker
- Arquitetura Hexagonal

## Iniciando desenvolvimento
- Clone o repositório
```bash
git clone <URL_DO_REPOSITORIO>
```
- Inicie os containers
```bash
docker-compose up -d
```
- Validando se o container do banco de dados está rodando
```bash
docker ps
```
- Acesse o container
```bash
docker exec -it appproduct bash
```
- Iniciando go mod (gerenciador de dependências do Go)
```bash
go mod init github.com/seu-usuario/seu-repositorio
```

- Criando um arquivo para o banco de dados sqlite
```bash
touch sqlite.db
```

- Acessando banco criado
```bash
sqlite3 sqlite.db
```

- Criando tabela
```bash
create table products(id string, name string, price float, status string);
```

- Consultando tabelas
```bash
.tables
```

## Unitários
### Gerando mocks para simular classes externas
- Acesse o container
```bash
docker exec -it appproduct bash
```
- Comando para geração dos mocks do módulo application e domínio product
```bash
mockgen -destination=application/mocks/application.go -source=application/product.go application
```