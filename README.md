# XGBoost prediction using Webassembly
This is a Proof of Concept of how a model-based prediction can be run within the browser via Webassembly.
Webassembly is a promising new technology allowing webapplications to run programs written in languages like Rust, C, C++, Java, or Golang and compiled into bytecode  directly in the Browser.

This repository serves a simple [web assembly](https://webassembly.org/) (wasm) application 
to perform a prediction based on a pre-trained xgb model, using data from a table in the browser, which can be loaded as a delimited file
by the user. We use the golang leaves  [leaves](https://gowalker.org/github.com/dmitryikh/leaves) package to do
the work. This is heavily inspired by  [this](https://vsoch.github.io/regression-wasm/) which uses golang and webassembly.

# Flow
- The xgboost model has been trained with python (see python_modeling).
- Once the model has been trained, it is exported as  binary file (*model.bst*).
- We then use the  [leaves](https://github.com/sajari/regression) golang package to compute predictions based on this model.
- We use the [syscall/js](https://golang.org/pkg/syscall/js/) package to interface the golang script with javascript
- We finally compile the golang script into Webassembly bytecode which can be run within the Browser.

As a matter of fact, intensive computing within the browser via Webassembly can be up to 20 times faster than the Javascript equivalent.


# Local

you need golang installed on your machine.

1 - build the wasm.

```bash
$ cd xgb-wasm
$ make
```

2 - Add your own Go version specific `wasm_exec.js` file :

```bash
$ cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./demo
```

3 - cd into the "demo" folder and execute `sheret.exe` (Windows users) to start a local server.

4 - open the browser at http://localhost:9999


# Credits
Vanessa Sochat / Lawrence Livermore National Laboratory /  vsoch.github.io
