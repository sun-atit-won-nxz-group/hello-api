name: CI Checks

on:
  push:
    branches:
      - main
  # pull_request: # ใส่ PR ไว้ด้วยเพื่อตรวจก่อน Merge
  #   branches:
  #     - main

jobs:
  # 1. Job สำหรับตรวจ Lint
  lint:
    name: Lint
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4 # อัปเดตเป็น v4 (ใหม่กว่า v2)
      - name: golangci-lint
        uses: golangci/golangci-lint-action@v3
        with:
          version: latest

  # 2. Job สำหรับ Test (จะรันก็ต่อเมื่อ lint ผ่านแล้ว)
  test:
    name: Test Application
    needs: lint # รอให้ lint ผ่านก่อนค่อยทำ
    runs-on: ubuntu-latest
    steps:
      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: '^1.18'

      - name: Check out code
        uses: actions/checkout@v4

      - name: Run Format Check
        run: make check-format

      - name: Vet
        run: go vet ./...

      - name: Run Test
        run: make test

      - name: Coverage Check
        run: make coverage

-------------------
linters-settings:
  stylecheck:
    checks: ["all", "ST1*"]

issues:
  exclude-use-default: false

  
 