<script setup>
import {reactive, onMounted} from 'vue'
import {DoAction, TestConnecting} from "../../wailsjs/go/api/EtcdApi"
import {List as ListOps} from "../../wailsjs/go/api/OperatorApi";
import {Get as GetGlobalConfig, Set as SetGlobalConfig} from "../../wailsjs/go/api/GlobalConfigApi";
import {JsonDiff, GetTempJson, TempStoreJson, GenHistoryName, GetHistory} from "../../wailsjs/go/api/TextApi";
import {ElMessage} from "element-plus";
import {Setting, Refresh, Collection, Tickets, DocumentChecked, Link, Loading} from "@element-plus/icons-vue";

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
  etcdEndPoint: "localhost:2379",
  testConnButtonClass: "info",
  testConnecting: false, // testConnecting 是否正在测试连接
  etcdEpSuggestions: [],

  resultText: "Please enter key below 👇",
  jsonContent: "默认占位内容",
  tips: "提示板",
  storeTmpReq: {
    name: "",
    tempJson: ""
  },
  tempJson: "",
  tempHistory: {},
  diffStatistic: {
    add: {},
    del: {},
    diff: {}
  },
  diff: {},


  doActionReq: {
    key: "",
    value: "",
    keyType: KTypeWholeKey,
    action: ATypeGet,
    doByOpHistory: false,
  },

  // 操作模块相关变量
  listOpsReq: {
    limit: 10,
  },
  opHistory: [],

  // ui controller
  settingPage: false,
  storeTmpJsonPage: false,
  storeConfirmDialog: false,
})

onMounted(() => {
  getGlobalConfig()
  listOps()
})

function getHistory() {
  GetHistory().then(result => {
    // 如果 result.code 不为 200，弹窗提示
    if (result.code !== 200) {
      ElMessage.error(result.message)
      return
    }
    console.log(result.data)
    data.tempHistory = result.data
  })
}

function openStoreTmpJsonConfirmDialog() {
  GenHistoryName(data.doActionReq.key).then(result => {
    // 如果 result.code 不为 200，弹窗提示
    if (result.code !== 200) {
      ElMessage.error(result.message)
      return
    }
    data.storeTmpReq.name = result.data
    data.storeConfirmDialog = true
  })
}

function getTmpJson() {
  GetTempJson().then(result => {
    // 如果 result.code 不为 200，弹窗提示
    if (result.code !== 200) {
      ElMessage.error(result.message)
      return
    }
    data.tempJson = result.data
  })
}

function storeTmpJson() {
  TempStoreJson(data.storeTmpReq.name, data.jsonContent).then(result => {
    // 如果 result.code 不为 200，弹窗提示
    if (result.code !== 200) {
      ElMessage.error(result.message)
      return
    }
    ElMessage.success(result.message)
    getTmpJson()
    data.storeConfirmDialog = false

    getHistory()
  })
}

function jsonDiff() {
  let req = {
    old: data.tempJson,
    new: data.jsonContent
  }
  JsonDiff(req).then(result => {
    // 如果 result.code 不为 200，弹窗提示
    if (result.code !== 200) {
      ElMessage.error(result.message)
      return
    }
    data.diffStatistic = result.data

    // 如果 diffStatistic 的三个字段都为空，说明没有差异，弹窗提示
    if (Object.keys(data.diffStatistic.add).length === 0 &&
        Object.keys(data.diffStatistic.del).length === 0 &&
        Object.keys(data.diffStatistic.diff).length === 0) {
      ElMessage.success("没有差异")
    }
    // data.diff = result.data.diff
  })
}

function getGlobalConfig() {
  GetGlobalConfig().then(result => {
    // 如果 result.code 不为 200，弹窗提示
    if (result.code !== 200) {
      ElMessage.error(result.message)
      return
    }
    data.jsonFmt = result.data.jsonFormat
    data.etcdEndPoint = result.data.etcdEndPoint
  })
}

function saveGlobalConfig() {
  let req = {
    jsonFormat: data.jsonFmt,
    etcdEndPoint: data.etcdEndPoint,
  }
  SetGlobalConfig(req).then(result => {
    // 如果 result.code 不为 200，弹窗提示
    if (result.code !== 200) {
      ElMessage.error(result.message)
      return
    }
    ElMessage.success(result.message)
    data.settingPage = false
    data.testConnButtonClass = "info"
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

    // 刷新操作历史
    listOps()
  })
}

// doActionByOpRecord 根据操作记录执行操作 （复现操作）
function doActionByOpRecord(row) {
  data.doActionReq = {
    key: row.key,
    value: row.value,
    keyType: row.keyType,
    action: row.action,
    doByOpHistory: true,
  }
  doAction()
  data.doActionReq.doByOpHistory = false
}

function listOps() {
  ListOps(data.listOpsReq).then(result => {
    // 如果 result.code 不为 200，弹窗提示
    if (result.code !== 200) {
      ElMessage.error(result.message)
      return
    }
    // todo: total
    data.opHistory = result.data
    console.log(result.data)
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
  // todo 取消日期的显示，只显示时间
  var datetime = date.toLocaleString();
  var time = datetime.split(" ")[1];
  return time; // 根据需要调整格式
}

function openSettingPage() {
  data.settingPage = true
  restaurants = loadAll()
}

var restaurants = [];

var querySearch = function(queryString, cb) {
  var results = queryString
      ? restaurants.filter(createFilter(queryString))
      : restaurants;
  // call callback function to return suggestions
  cb(results);
};

var createFilter = function(queryString) {
  return function(restaurant) {
    return (
        restaurant.value.toLowerCase().indexOf(queryString.toLowerCase()) === 0
    );
  };
};

var loadAll = function() {
  return [
    { value: 'localhost:2379', link: '' },
    { value: '10.16.207.141:2379', link: '' },
    { value: '10.16.204.117:2379', link: '' },
  ];
};

function clickEtcdEpSuggestion(item) {
  testConnecting()
}

function testConnecting() {
  data.testConnecting = true
  TestConnecting(data.etcdEndPoint).then(result => {
    // 如果 result.code 不为 200，弹窗提示
    if (result.code !== 200) {
      ElMessage.error(result.message)
      data.testConnButtonClass = "danger"
      data.testConnecting = false
      return
    }
    ElMessage.success(result.message)
    data.testConnButtonClass = "success"
    data.testConnecting = false
  })
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
            <el-tooltip class="box-item" effect="dark" content="根据 key 获取值" placement="top-start">
              <el-button size="small" @click="get" type="primary" plain>Get</el-button>
            </el-tooltip>
            <el-tooltip class="box-item" effect="dark" content="更新 key 的值" placement="top-start">
              <el-button size="small" @click="put" type="primary" plain>Put</el-button>
            </el-tooltip>
            <el-tooltip class="box-item" effect="dark" content="删除" placement="top-start">
              <el-button size="small" @click="del" type="primary" plain>Del</el-button>
            </el-tooltip>

            <el-tooltip class="box-item" effect="dark" content="根据前缀获取 key 列表" placement="top-start">
              <el-button size="small" type="success" plain @click="listKeyByPfx">listKey</el-button>
            </el-tooltip>
            <el-tooltip class="box-item" effect="dark" content="根据前缀获取 值 列表" placement="top-start">
              <el-button size="small" type="success" plain @click="listValByPfx">listVal</el-button>
            </el-tooltip>
            <el-tooltip class="box-item" effect="dark" content="根据前缀删除" placement="top-start">
              <el-button size="small" type="success" plain @click="delByPfx">del</el-button>
            </el-tooltip>

            <el-tooltip class="box-item" effect="dark" content="根据关键字获取 key 列表" placement="top-start">
              <el-button size="small" type="warning" plain @click="listKeyByKw">listKey</el-button>
            </el-tooltip>
            <el-tooltip class="box-item" effect="dark" content="根据关键字获取 值 列表" placement="top-start">
              <el-button size="small" type="warning" plain @click="listValByKw">listVal</el-button>
            </el-tooltip>
            <el-tooltip class="box-item" effect="dark" content="根据关键字删除" placement="top-start">
              <el-button size="small" type="warning" plain @click="delByKw">del</el-button>
            </el-tooltip>

            <el-button type="info" size="small" :icon="Setting" circle @click="openSettingPage"/>
            <el-tooltip class="box-item" effect="dark" content="更新操作历史列表" placement="top-start">
              <el-button type="info" size="small" :icon="Refresh" circle @click="listOps"/>
            </el-tooltip>
          </div>
          <div id="text" style="margin-top: 10px">
            <el-popover :visible="data.storeConfirmDialog" placement="top" :width="400">
                <el-input
                    v-model="data.storeTmpReq.name"
                    placeholder="Please Input"
                    style="margin-bottom: 10px"
                    size="small"
                />
              <div style="text-align: right; margin: 0">
                <el-button size="small" text @click="data.storeConfirmDialog = false">取消</el-button>
                <el-button type="info" text @click="storeTmpJson">
                  确定
                </el-button>
              </div>
              <template #reference>
                <el-tooltip class="box-item" effect="dark" content="暂存当前值" placement="top-start">
                  <el-button type="info" :icon="DocumentChecked" circle size="small" @click="openStoreTmpJsonConfirmDialog"/>
                </el-tooltip>
              </template>
            </el-popover>
            <el-tooltip class="box-item" effect="dark" content="比较左右两边的差异(暂不支持嵌套结构体)" placement="top-start">
              <el-button type="info" :icon="Tickets" circle size="small" @click="jsonDiff" />
            </el-tooltip>

            <el-popover placement="right" :width="300" trigger="click">
              <template #reference>
                <el-button type="info" size="small" :icon="Collection" circle @click="getHistory"/>
              </template>
              <div v-for="(tmpVal, name) in data.tempHistory" :key="name">
                <el-button
                    :key="name"
                    type="info"
                    text
                    size="small"
                    @click="data.tempJson = tmpVal"
                >{{ name }}</el-button>
              </div>
            </el-popover>
          </div>
        </div>
      </el-col>
      <el-col :span="12">
        <div class="tips">
          <el-table
              :data="data.opHistory"
              style="width: 100%"
              height="250"
              :row-class-name="opHistoryTableRowClassName">
            <el-table-column fixed label="op" width="60">
              <template #default="scope">
                <el-button
                    link
                    type="primary"
                    size="small"
                    @click="doActionByOpRecord(scope.row)"
                >执行</el-button>
              </template>
            </el-table-column>
            <el-table-column label="Id" prop="id" fixed width="50px"></el-table-column>
            <el-table-column label="描述" prop="desc" width="400px"></el-table-column>
            <el-table-column label="Result" prop="result">
              <template #default="{row}">
                <el-tag
                    :type="row.result === 0 ? 'success' : 'danger'"
                    disable-transitions
                >{{ row.result === 0 ? "成功" : "失败" }}</el-tag
                >
              </template>
            </el-table-column>
            <el-table-column label="CreateAt" prop="createAt">
              <template #default="{row}">
                {{ formatDate(row.createAt) }}
              </template>
            </el-table-column>
          </el-table>
        </div>
      </el-col>
    </el-row>

    <el-row :gutter="20">
      <el-col :span="12">
        <el-input
            type="textarea"
            style="margin-top: 10px"
            :rows="20"
            placeholder="请输入内容"
            v-model="data.jsonContent">
        </el-input>
      </el-col>
      <el-col :span="12">
        <el-input
            type="textarea"
            style="margin-top: 10px"
            :rows="20"
            placeholder="请输入内容"
            :disabled="true"
            v-model="data.tempJson">
        </el-input>
      </el-col>
    </el-row>

    <el-row :gutter="20">
      <el-col :span="12">
        <!--  当  diffStatistic 不为空时，遍历，并使用标签来展示   -->
        <div id="diff-statistic" style="margin-top: 10px; text-align: left" >
          <div v-for="(changes, fieldName) in data.diffStatistic.diff" :key="fieldName">
            <el-tag type="warning">{{ fieldName }}</el-tag>
            <el-tag type="info" class="diff-value-item">"{{ changes.old }}" -> "{{ changes.new }}"</el-tag>
          </div>
          <div v-for="(value, fieldName) in data.diffStatistic.add" :key="fieldName">
            <el-tag type="success">{{ fieldName }}</el-tag>
            <el-tag type="info" class="diff-value-item">{{ value }}</el-tag>
          </div>
          <div v-for="(value, fieldName) in data.diffStatistic.del" :key="fieldName">
            <el-tag type="danger">{{ fieldName }}</el-tag>
            <el-tag type="info" class="diff-value-item">{{ value }}</el-tag>
          </div>
        </div>
      </el-col>
      <el-col :span="12">

      </el-col>
    </el-row>

    <el-dialog
        v-model="data.storeTmpJsonPage"
        title="暂存当前的值"
        width="30%"
    >
      <template #footer>
      <span class="dialog-footer">
        <el-input
            v-model="data.storeTmpReq.name"
            placeholder="Please Input"
            style="margin-bottom: 10px"
        />
        <el-button @click="data.storeTmpJsonPage = false">Cancel</el-button>
        <el-button type="primary" @click="storeTmpJson">
          Confirm
        </el-button>
      </span>
      </template>
    </el-dialog>

    <!--  全局配置  -->
    <el-drawer
        v-model="data.settingPage"
        title="global setting"
    >
      <el-form label-width="120px">
        <el-form-item>
          <el-autocomplete
              v-model="data.etcdEndPoint"
              :fetch-suggestions="querySearch"
              clearable
              class="inline-input w-50"
              placeholder="Etcd EndPoint"
              @select="clickEtcdEpSuggestion"
          />
          <el-tooltip effect="dark" content="测试连接" placement="top-start">
            <el-button v-if="!data.testConnecting" :type="data.testConnButtonClass" size="small" :icon="Link" circle @click="testConnecting" style="margin-left: 5px"/>
            <el-button v-if="data.testConnecting" type="info" size="small" :icon="Loading" circle style="margin-left: 5px"/>
          </el-tooltip>
        </el-form-item>
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

.diff-statistic-item {
  display: flex;
  align-items: flex-start;
  margin-bottom: 2px
}

.diff-value-item {
  margin-top: 2px;
  margin-left: 2px;
}

.tooltip-base-box .box-item {
  width: 110px;
  margin-top: 10px;
}

</style>
