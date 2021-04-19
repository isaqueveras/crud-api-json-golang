# Criando uma simples API usando a linguagem Golang
Neste repositorio, contém os arquivos que desenvolvi estudando a linguagem Go para fazer uma API que retorna dados em JSON.

Criando um servidor de HTTP em Golang que serve dados em formato JSON. Essa API REST em Go, com operações de cadastrar, listar, modificar, e deletar recursos (o chamado CRUD). Estou começando do zero, aprendendendo os meus primeiros passos de desenvolvimento Web com a linguagem Golang.

### Por onde estou estudando?
Estou estudando para contruir essa API atravez de uma playlist do canal _NBK Mundo Tech_ onde estou dando meus primeiros passos nesta linguagem que é usado por muita gente.

#### 1° Aula
Na primeira aula, preparei a base da aplicação. Usando o _http_ atravez do import _"net/http"_. Também aprendi a imprimir na tela usando a rota _"/"_.

### 2° Aula
Vamos fazer a refatorização do nosso código para uma organização melhor, fazendo com que cada funcionalidade se torne uma função.

Usando _fmt.Println_ temos a resposta via terminal de que o servidor esta funcionando corretamente.
```go
fmt.Println("Servidor rodando na porta 1317")
```
Fiz a refatoração do codigo, criei mais 3 funções que cada uma tem suas responsabilidades.

```go
// Função que funciona como um Controller da aplicação 
func rotaPrincipal(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Olá mundo!")
}

// Função de administração das rotas
func configurarRotas() {
  http.HandleFunc("/", rotaPrincipal)
}

// Essa função recebe todas as configurações do servidor
func configurarServidor() {
  configurarRotas()

  // Escrevendo no terminal
  fmt.Println("Servidor rodando na porta 1317")
  
  // Adicionando a porta do servidor
  http.ListenAndServe(":1337", nil) 
}
```
E a função _main_ ficou bem menor com as alterações, assim, facilitando o entendimento do códico.

```go
func main() {
  configurarServidor()
}
```