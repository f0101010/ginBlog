<template>
    <div class="container">
        <div class="loginBox">
            <a-form :rules="rules" :model="formdata" class="loginForm">
                <a-form-item name="username">
                    <a-input v-model:value="formdata.username" placeholder="Username"> </a-input>
                </a-form-item>
                <a-form-item name="password">
                    <a-input-password v-model:value="formdata.password" placeholder="Password"></a-input-password>
                </a-form-item>
                <a-form-item class="loginBtn">
                    <a-button type="primary">登录</a-button>
                </a-form-item>
            </a-form>
        </div>
    </div>
</template>


<script>
export default {
    data() {
        return {
            formdata: {
                username: '',
                password: '',
            },
            rules: {
                username: [
                    { required: true, message: '请输入用户名', trigger: 'blur' },
                    {
                        min: 4,
                        max: 12,
                        message: '用户名必须在4到12个字符之间',
                        trigger: 'blur',
                    },
                ],
                password: [
                    { required: true, message: '请输入密码', trigger: 'blur' },
                    {
                        min: 6,
                        max: 20,
                        message: '密码必须在6到20个字符之间',
                        trigger: 'blur',
                    },
                ]
            },
        }
    },
}
</script>


<style scoped>
.container {
    height: 100%;
    background-color: #282C34;
}

.loginBox {
    width: 450px;
    height: 300px;
    background-color: #fff;
    position: absolute;
    top: 50%;
    left: 70%;
    transform: translate(-50%, -50%);
    border-radius: 9px;
}

.loginForm {
    width: 100%;
    position: absolute;
    bottom: 10%;
    padding: 0 20px;
    box-sizing: border-box;
}

.loginBtn {
    display: flex;
    justify-content: flex-end;
}
</style>