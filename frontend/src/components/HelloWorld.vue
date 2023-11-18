<script setup>
import {reactive} from 'vue'
import {Greet} from '../../wailsjs/go/main/App'
import {Get2String, PutCompact, Del, ListKeyByPrefix, ListKeyByKeyword, ListValueByKeyword} from '../../wailsjs/go/etcd/EtcdClient'

const data = reactive({
  name: "",
  resultText: "Please enter key below 👇",
  jsonContent: "默认占位内容",
  tips: "提示板"
})

function greet() {
  const person = {
    name: data.name,
    nickName: "索隆不喝酒"
  }
  Greet(person).then(result => {
    data.resultText = result
  })
}

function get() {
  Get2String(data.name).then(result => {
    data.jsonContent = result
  })
}

function put() {
  PutCompact(data.name, data.jsonContent).then(result => {

  })
}

function del() {
  Del(data.name).then(result => {

  })
}

function listKeyByPfx() {
  ListKeyByPrefix(data.name).then(result => {
    data.jsonContent = result
  })
}

function listKeyByKw() {
  ListKeyByKeyword(data.name).then(result => {
    data.jsonContent = result
  })
}

function listValByKw() {
  ListValueByKeyword(data.name).then(result => {
    data.jsonContent = result
  })
}

</script>

<template>
  <main>
    <!--  将上面这一部分分为左右两边，把现在的内容都放到左边，右边预留  -->
    <el-row :gutter="20">
      <el-col :span="16">
        <div class="ops">
          <el-input v-model="data.name" placeholder="请输入「key」「关键字」「前缀」任一种" class="my-input"></el-input>
          <el-button size="mini" @click="get" type="primary" plain>Get</el-button>
          <el-button size="mini" @click="put" type="primary" plain>Put</el-button>
          <el-button size="mini" @click="del" type="danger" plain>Del</el-button>
          <el-button size="mini" @click="listValByKw">listValByKw</el-button>
          <el-button size="mini" @click="listKeyByPfx">listByPfx</el-button>
          <el-button size="mini" @click="listKeyByKw">listByKw</el-button>
        </div>
      </el-col>
      <el-col :span="8">
        <div class="tips">
          <el-input
              type="textarea"
              :autosize="{ minRows: 3, maxRows: 5}"
              placeholder="请输入内容"
              v-model="data.tips"
              :disabled="true"
          >
          </el-input>
        </div>
      </el-col>
    </el-row>

    <el-input
        type="textarea"
        style="margin-top: 10px"
        :autosize="{ minRows: 2, maxRows: 30}"
        placeholder="请输入内容"
        v-model="data.jsonContent">
    </el-input>
  </main>
</template>

<style scoped>
.my-input {
  margin-bottom: 10px;
}

</style>
