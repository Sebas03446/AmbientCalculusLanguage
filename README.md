# AmbientCalculusLanguage

## 1. Introduction

This project provides a comprehensive implementation of the Ambient Calculus programming language. It covers the essential aspects of running the language, grammatical categories, code descriptions, and concluding remarks. Notably, the project leverages ANTLR4 with Go and Dot to facilitate syntax tree generation.

![ast](https://github.com/user-attachments/assets/e0484745-bfc7-4206-9118-c379f2ed3894)


## 2. Steps to Execute the Language

1. **Clone the Repository:**
   ```bash
   git clone [https://github.com/Sebas03446/AmbientCalculusLanguage.git](https://github.com/Sebas03446/AmbientCalculusLanguage.git)
   ```
2. **Execute ANTLR4**
  ```bash
    curl -O [https://www.antlr.org/download/antlr-4.9.2-complete.jar](https://www.antlr.org/download/antlr-4.9.2-complete.jar)
    export CLASSPATH=".:antlr-4.9.2-complete.jar:$CLASSPATH"
    alias antlr4='java -jar antlr-4.9.2-complete.jar'
    alias grun='java org.antlr.v4.gui.TestRig'
  ```
3. **Update the parser**
  ```bash
     antlr4 -Dlanguage=Go AmbientCalculus.g4 -visitor -o parser
  ```
4. **Run the Compilator**
  ```bash
    go run main.go ejemplo1.ambient
  ```

This will compile and execute your code, generating the AST representation and its visualization as an image.
