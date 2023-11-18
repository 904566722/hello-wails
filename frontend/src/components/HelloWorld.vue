<script setup>
import {reactive} from 'vue'
import {Get2String, PutCompact, Del, ListKeyByPrefix, ListKeyByKeyword, ListValueByKeyword} from '../../wailsjs/go/etcd/EtcdClient'
import {DoAction} from "../../wailsjs/go/api/EtcdApi"
import {List as ListOps} from "../../wailsjs/go/api/OperatorApi";
import {ElMessage} from "element-plus";

const KTypeWholeKey = 0
const KTypePrefix = 1
const KTypeKeyword = 2

const ATypeGet = 0
const ATypePut = 1
const ATypeDel = 2
const ATypeListKey = 3
const ATypeListVal = 4
const ATypeList = 5

const data = reactive({
  name: "",
  keyType: KTypeWholeKey,
  action: ATypeGet,

  resultText: "Please enter key below 👇",
  jsonContent: "默认占位内容",
  tips: "提示板",


  doActionReq: {
    data: "",
    keyType: KTypeWholeKey,
    action: ATypeGet,
  },

  // 操作模块相关变量
  listOpsReq: {
    limit: 10,
  },
  opHistory: [],
})

// 刷新请求内容
// todo: delete code
function freshReqContent() {
  data.doActionReq = {
    data: data.name,
    keyType: data.keyType,
    action: data.action,
  }
}

function doAction() {
  DoAction(data.doActionReq).then(result => {
    // 如果 result.code 不为 200，弹窗提示
    if (result.code !== 200) {
      ElMessage.error(result.message)
      return
    }
    data.jsonContent = result.data
    ElMessage.success(result.message)
  })
}

function listOps() {
  ListOps(data.listOpsReq).then(result => {
    console.log(result)
    // 如果 result.code 不为 200，弹窗提示
    if (result.code !== 200) {
      ElMessage.error(result.message)
      return
    }
    // todo: total
    data.opHistory = result.data
    ElMessage.success(result.message)
  })
}

function get() {
  data.doActionReq = {
    data: data.name,
    keyType: KTypeWholeKey,
    action: ATypeGet,
  }

  doAction()
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

function opHistoryTableRowClassName({row, rowIndex}) {
  if (row.result === 0) {
    return 'success-row';
  } else {
    return 'warning-row';
  }
}

function formatDate(dateString) {
  const date = new Date(dateString);
  // 这里可以使用你喜欢的日期格式化库，比如 moment.js 或 date-fns
  // 在这个简单的示例中，使用 JavaScript 原生的方法
  return date.toLocaleString(); // 根据需要调整格式
}

</script>

<template>
  <main>
    <!--  将上面这一部分分为左右两边，把现在的内容都放到左边，右边预留  -->
    <el-row :gutter="20">
      <el-col :span="12">
        <div class="ops">
          <el-input v-model="data.name" placeholder="请输入「key」「关键字」「前缀」任一种" class="my-input"></el-input>
          <el-button size="small" @click="get" type="primary" plain>Get</el-button>
          <el-button size="small" @click="put" type="primary" plain>Put</el-button>
          <el-button size="small" @click="del" type="danger" plain>Del</el-button>
          <el-button size="small" @click="listValByKw">listValByKw</el-button>
          <el-button size="small" @click="listKeyByPfx">listByPfx</el-button>
          <el-button size="small" @click="listKeyByKw">listByKw</el-button>
        </div>
      </el-col>
      <el-col :span="12">
        <div class="tips">
<!--          <el-input-->
<!--              type="textarea"-->
<!--              :autosize="{ minRows: 3, maxRows: 5}"-->
<!--              placeholder="请输入内容"-->
<!--              v-model="data.tips"-->
<!--              :disabled="true"-->
<!--          >-->
<!--          </el-input>-->
          <el-table
              :data="data.opHistory"
              style="width: 100%"
              height="250"
              :row-class-name="opHistoryTableRowClassName">
            <el-table-column label="Id" prop="id" fixed></el-table-column>
            <el-table-column label="KeyType" prop="keyType"></el-table-column>
            <el-table-column label="Action" prop="action"></el-table-column>
            <el-table-column label="Result" prop="result"></el-table-column>
            <el-table-column label="Message" prop="message"></el-table-column>
            <el-table-column label="CreateAt" prop="createAt">
              <template #default="{row}">
                {{ formatDate(row.createAt) }}
              </template>
            </el-table-column>
            <el-table-column fixed="right" label="op" width="60">
              <template #default>
                <el-button link type="primary" size="small">执行</el-button>
              </template>
            </el-table-column>
          </el-table>


          <el-button size="small" @click="listOps">listOpHistory</el-button>
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

.el-table .warning-row {
  background: #d53d72;
}

.el-table .success-row {
  background: #29c731;
}

</style>
