<template>
  <div class="container">
    <video v-if="false" class="container-bg" src="" autoplay="autoplay" loop="loop" muted="muted"></video>
    <div v-if="!defalut_data.showRegisterForm" class="container-form">
      <form class="cus-form">
        <h2 class="form-title">登录</h2>
        <div class="cus-form-item">
          <input type="text" v-model="data.userName" placeholder="手机号" class="cus-input">
          <i class="cus-item-icon iconfont icon-account"/>
        </div>
        <div class="cus-form-item">
          <i class="cus-item-icon iconfont icon-password"/>
          <input :type="`${data.pwdIsShow ? 'text' : 'password'}`" v-model="data.password" placeholder="密码" class="cus-input" @keydown.enter="login">
          <i :class="`cus-pwd iconfont ${data.pwdIsShow ? 'icon-eye2' : 'icon-eye'}`" @click="data.pwdIsShow = !data.pwdIsShow" />
        </div>
        <el-button native-type="button" type="primary" size="large" color="#9b49e7" round @click="login">登录</el-button>
        <el-button native-type="button" type="primary" size="large" color="#9b49e7" round @click="defalut_data.showRegisterForm = true">注册</el-button>
      </form>
    </div>
    <div v-else class="container-form">
      <!-- Register form goes here -->
      <form class="cus-form">
        <h2 class="form-title">注册</h2>
        <div class="cus-form-item">
          <input type="text" v-model="registerData.userName" placeholder="请填写手机号" class="cus-input">
          <i class="cus-item-icon iconfont icon-account"/>
        </div>
        <div class="cus-form-item">
          <i class="cus-item-icon iconfont icon-password"/>
          <input :type="`${registerData.pwdIsShow ? 'text' : 'password'}`" v-model="registerData.password" placeholder="请输入密码" class="cus-input" @keydown.enter="register">
          <i :class="`cus-pwd iconfont ${registerData.pwdIsShow ? 'icon-eye2' : 'icon-eye'}`" @click="registerData.pwdIsShow = !registerData.pwdIsShow" />
        </div>
        <el-button native-type="button" type="primary" size="large" color="#9b49e7" round @click="register">注册</el-button>
        <el-button native-type="button" type="primary" size="large" color="#9b49e7" round @click="defalut_data.showRegisterForm = false">返回登录</el-button>
      </form>
    </div>
  </div>
</template>
  
  <script setup lang="ts">
  
  import {reactive} from "vue";
  import {ApiPost} from "../utils/request";
  import {ElMessage} from "element-plus";
  import {useRouter} from "vue-router";
  
  const data = reactive({
    userName: "",
    password: "",
    pwdIsShow: false,
  })
const defalut_data = reactive({
  showRegisterForm: false,
  phoneNumber:'',
  message:'请输入正确的手机号！'
})
  const registerData = reactive({
    userName: "",
    password: "",
    pwdIsShow: false,
  })
  
  const router = useRouter()
  const login = ()=>{
    ApiPost('/login', {userName: data.userName, password: data.password}).then((resp:any)=>{
      if (resp.code == 0) {
        router.push('/manage/index')
      } else {
        ElMessage.error({message: resp.msg})
      }
    })
  }
  
  function validatePhoneNumber(phoneNumber) {
      const phoneRegex = /^1[3456789]\d{9}$/; // 手机号正则表达式

      if (phoneRegex.test(phoneNumber)) {
        return true; // 格式正确
      }else if (phoneNumber.length !==11){
        return false; // 格式不正确
      } 
      else if(''){
        return false; // 格式不正确
      }
    }
   function register() {
      // 校验手机号格式
      if (!validatePhoneNumber(registerData.userName)) {
        ElMessage.error({message: defalut_data.message})
        return;
      }}


  </script>
  
  
  <style scoped>
  .container {
    width: 100vw;
    height: 100vh;
    color: #343333;
    background: url("../assets/image/user_login_backgroup.gif");
  }
  .container-bg {
    background-color: #b07ada;
    width: 100%;
    height: 100%;
    object-fit: cover;
    position: absolute;
    top: 0;
    left: 0;
  }
  
  .container-form {
    background: rgba(255,255,255,0.45);
    width: 480px;
    height: 460px;
    border-radius: 8px;
    position: relative;
    top: 20%;
    left: 30%;
  }
  .form-title{
    color: #6e00bf;
  }
  
  .cus-input {
    font-size: 16px;
    width: 100%;
    padding: 0  40px;
    border: none;
    outline-style: none ;
    border-radius: 26px;
    min-height: 40px;
    background: rgba(255,255,255,0.55);
  }
  .cus-input:focus{
      outline: 2px solid rgb( 169,52,217,72%);
      border: 0;
  }
  .cus-form {
    display: flex;
    flex-direction: column;
    padding: 10px 8px;
    gap: 32px;
  }
  .el-button {
    width: 70%;
    margin: 0 auto!important;
  }
  
  /*自定义密码框组件*/
  .cus-form-item {
    margin: 0 auto;
    width: 80%;
    position: relative;
  }
  .cus-item-icon{
    position: absolute;
    left: 18px;
    top: 6px;
    color: #b07ada;
  }
  .cus-pwd {
    color: #b07ada;
    position: absolute;
    right: 20px;
    top: 6px;
    cursor: pointer;
  }
  </style>