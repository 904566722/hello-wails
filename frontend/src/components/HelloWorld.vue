<script setup>
import {reactive, onMounted} from 'vue'
import {Get2String, PutCompact, Del, ListKeyByPrefix, ListKeyByKeyword, ListValueByKeyword} from '../../wailsjs/go/etcd/EtcdClient'
import {DoAction} from "../../wailsjs/go/api/EtcdApi"
import {List as ListOps} from "../../wailsjs/go/api/OperatorApi";
import {Get as GetGlobalConfig, Set as SetGlobalConfig} from "../../wailsjs/go/api/GlobalConfigApi";
import {ElMessage} from "element-plus";
import {Setting} from "@element-plus/icons-vue";

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

  jsonFmt: true,

  resultText: "Please enter key below 👇",
  jsonContent: "默认占位内容",
  tips: "提示板",


  doActionReq: {
    key: "",
    value: "",
    keyType: KTypeWholeKey,
    action: ATypeGet,
  },

  // 操作模块相关变量
  listOpsReq: {
    limit: 10,
  },
  opHistory: [],

  // ui controller
  settingPage: false

})

onMounted(() => {
  getGlobalConfig()
})


function getGlobalConfig() {
  GetGlobalConfig().then(result => {
    // 如果 result.code 不为 200，弹窗提示
    if (result.code !== 200) {
      ElMessage.error(result.message)
      return
    }
    data.jsonFmt = result.data.jsonFormat
  })
}

function saveGlobalConfig() {
  let req = {
    jsonFormat: data.jsonFmt,
  }
  SetGlobalConfig(req).then(result => {
    // 如果 result.code 不为 200，弹窗提示
    if (result.code !== 200) {
      ElMessage.error(result.message)
      return
    }
    ElMessage.success(result.message)
  })
}

function doAction() {
  DoAction(data.doActionReq).then(result => {
    // 如果 result.code 不为 200，弹窗提示
    if (result.code !== 200) {
      ElMessage.error(result.message)
      return
    }
    console.log(result.data)
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

// get 根据完整的 key 获取 value
function get() {
  data.doActionReq.keyType = KTypeWholeKey
  data.doActionReq.action = ATypeGet
  doAction()
}

// put 根据完整的 key 更新 value
function put() {
  data.doActionReq.keyType = KTypeWholeKey
  data.doActionReq.action = ATypePut
  data.doActionReq.value = data.jsonContent
  doAction()
}

// del 根据完整的 key 删除 value
function del() {
  data.doActionReq.keyType = KTypeWholeKey
  data.doActionReq.action = ATypeDel
  doAction()
}

// listKeyByPfx 根据前缀获取 key 列表
function listKeyByPfx() {
  data.doActionReq.keyType = KTypePrefix
  data.doActionReq.action = ATypeListKey
  doAction()
}

// listValByPfx 根据前缀获取 value 列表
function listValByPfx() {
  data.doActionReq.keyType = KTypePrefix
  data.doActionReq.action = ATypeListVal
  doAction()
}

// delByPfx 根据前缀删除 key-value
function delByPfx() {
  data.doActionReq.keyType = KTypePrefix
  data.doActionReq.action = ATypeDel
  doAction()
}

// listKeyByKw 根据关键字获取 key 列表
function listKeyByKw() {
  data.doActionReq.keyType = KTypeKeyword
  data.doActionReq.action = ATypeListKey
  doAction()
}

// listValByKw 根据关键字获取 value 列表
function listValByKw() {
  data.doActionReq.keyType = KTypeKeyword
  data.doActionReq.action = ATypeListVal
  doAction()
}

// delByKw 根据关键字删除 key-value
function delByKw() {
  data.doActionReq.keyType = KTypeKeyword
  data.doActionReq.action = ATypeDel
  doAction()
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
          <el-input v-model="data.doActionReq.key" placeholder="请输入「key」「关键字」「前缀」任一种" class="my-input"></el-input>
          <div id="whole-key">
            <el-button size="small" @click="get" type="primary" plain>Get</el-button>
            <el-button size="small" @click="put" type="primary" plain>Put</el-button>
            <el-button size="small" @click="del" type="danger" plain>Del</el-button>
          </div>
          <div id="prefix-key">
            <el-button size="small" @click="listKeyByPfx">listKey</el-button>
            <el-button size="small" @click="listValByPfx">listVal</el-button>
            <el-button size="small" @click="delByPfx">del</el-button>
          </div>
          <div id="keyword-key">
            <el-button size="small" @click="listKeyByKw">listKey</el-button>
            <el-button size="small" @click="listValByKw">listVal</el-button>
            <el-button size="small" @click="delByKw">del</el-button>
          </div>
        </div>
        <el-switch
            v-model="data.jsonFmt"
            class="mb-2"
            inactive-text="压缩"
            active-text="格式化"
        />
        <el-button type="info" :icon="Setting" circle @click="data.settingPage = true"/>
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
            <el-table-column label="描述" prop="desc"></el-table-column>
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

    <!--  全局配置  -->
    <el-drawer
        v-model="data.settingPage"
        title="global setting"
    >
      <el-form label-width="120px">
        <el-form-item label="json format">
          <el-switch
              v-model="data.jsonFmt"
              inactive-text="压缩"
              active-text="格式化"
          />
        </el-form-item>
      </el-form>

      <el-button type="primary" plain @click="saveGlobalConfig">Save</el-button>
    </el-drawer>
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
