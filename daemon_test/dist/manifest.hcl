name = "daemon_test"
authors = ["OptimusePrime"]
version = "1.0.0"

frontend {
  tab_name = "File Manager"
  platform = "svelte"
}

backend {
  commands = ["test_cmd"]
  main = "./daemon_test.wasm"
}