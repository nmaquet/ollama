//go:build linux && nvidia
// +build linux,nvidia

package llm

//go:generate git submodule init
//go:generate git submodule update --force ggml
//go:generate git -C ggml apply ../ggml_patch/0001-add-detokenize-endpoint.patch
//go:generate git -C ggml apply ../ggml_patch/0002-34B-model-support.patch
//go:generate cmake -S ggml -B ggml/build/gpu -DLLAMA_CUBLAS=ON -DLLAMA_K_QUANTS=on
//go:generate cmake --build ggml/build/gpu --target server --config Release
