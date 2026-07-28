---
name: commit-message
description: Gera mensagens de commit curtas em português no padrão conventional commits (tipo: descrição), analisando o diff em stage. Use sempre que o usuário pedir uma mensagem de commit, perguntar "o que escrevo nesse commit", pedir para descrever/resumir mudanças para commitar, ou usar /commit-message — mesmo que não cite "conventional commits" explicitamente.
---

# Mensagem de commit

Gera a mensagem e a entrega pronta para o usuário copiar. **Não execute `git commit`** — o usuário revisa e commita por conta própria.

## Processo

1. Leia o que será commitado:
   ```bash
   git status --short && git diff --staged --stat && git diff --staged
   ```
2. Se não houver nada em stage, olhe `git diff` (não-staged) e avise o usuário que nada está em stage antes de dar a mensagem.
3. Se o diff for muito grande, use `git diff --staged --stat` para o panorama e leia por inteiro só os arquivos que definem a mudança.
4. Escreva a mensagem no formato abaixo e apresente em um bloco de código, sozinha, para facilitar a cópia.

## Formato

```
tipo: descrição em português, minúscula, sem ponto final
```

Uma linha só, no máximo ~72 caracteres. Sem escopo entre parênteses, sem corpo, sem rodapé — o histórico deste projeto é assim e a consistência importa mais do que detalhamento.

Verbo no presente da terceira pessoa, descrevendo o que o commit faz no código: "adiciona", "corrige", "remove", "atualiza", "renomeia", "move", "extrai".

### Tipos

| Tipo | Quando usar |
|------|-------------|
| `feat` | funcionalidade nova ou mudança de comportamento visível |
| `fix` | correção de bug |
| `refactor` | reestrutura código sem mudar comportamento |
| `chore` | dependências, configs, arquivos auxiliares |
| `ci` | workflows, pipelines, Docker de build |
| `docs` | documentação, README, comentários |
| `test` | testes |
| `style` | formatação, imports, sem efeito em lógica |

Na dúvida entre dois tipos, pergunte-se o que o leitor do histórico procuraria daqui a seis meses. Uma mudança que conserta algo quebrado é `fix` mesmo que adicione linhas; um endpoint novo é `feat` mesmo que seja pequeno.

## O que faz uma boa descrição

Diga **o que mudou no comportamento do sistema**, não quais arquivos foram tocados. `git show` já lista os arquivos; a mensagem existe para poupar essa leitura.

Evite genéricos ("ajustes", "melhorias", "diversas correções") — eles ocupam espaço sem informar. Se a mudança é realmente pequena e mecânica, seja específico sobre ela mesmo assim.

**Exemplos:**

Diff: novo handler + rota `POST /api/url` que cria URLs curtas
→ `feat: adiciona endpoint de criação de url encurtada`

Diff: `generateRandomString` gerava códigos de 11 chars por erro de índice
→ `fix: corrige tamanho do código gerado no encurtamento`

Diff: lógica de query movida de `service` para `repository`
→ `refactor: move acesso a dados do service para o repository`

Diff: `go.mod`/`go.sum` com bump de versão do chi
→ `chore: atualiza chi para v5.2.0`

Ruim: `feat: atualiza arquivos` · `fix: correções no service` · `feat: várias melhorias`

## Vários assuntos no mesmo diff

Quando o stage mistura mudanças não relacionadas, uma linha só não descreve bem nenhuma delas. Diga isso ao usuário e sugira como dividir em commits — indicando quais arquivos vão em cada um e a mensagem de cada — em vez de forçar uma descrição guarda-chuva. Se ele preferir um commit único, escolha o tipo da mudança dominante e descreva a parte mais relevante.
