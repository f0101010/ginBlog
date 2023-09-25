<template>
    <div class="container">
        <div class="loginBox">
            <a-form ref="loginFormRef" :rules="rules" :model="formdata" class="loginForm">
                <a-form-item name="username">
                    <a-input v-model:value="formdata.username" placeholder="Username">
                    </a-input>
                </a-form-item>
                <a-form-item name="password">
                    <a-input-password v-model:value="formdata.password" placeholder="Password"></a-input-password>
                </a-form-item>
                <a-form-item class="loginBtn">
                    <a-button type="primary" @click="login">登录</a-button>
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
                    {
                        required: true,
                        message: '请输入用户名',
                        trigger: 'blur',
                    },
                    {
                        min: 4,
                        max: 12,
                        message: '用户名必须在4到12个字符之间',
                        trigger: 'blur',
                    },
                ],
                password: [
                    {
                        required: true,
                        message: '请输入密码',
                        trigger: 'blur',
                    },
                    {
                        min: 6,
                        max: 20,
                        message: '密码必须在6到20个字符之间',
                        trigger: 'blur',
                    },
                ],
            }
        }
    },
    methods: {
        async login() {
            try {
                // 手动触发表单验证，等待验证结果
                const valid = await this.$refs.loginFormRef.validate();

                if (!valid) {
                    this.$message.error('输入非法数据，请重新输入');
                    return;
                }

                // 如果验证通过，执行登录请求
                const { data: res } = await this.$http.post('login', this.formdata);

                if (res.status !== 200) {
                    this.$message.error(res.message);
                } else {
                    // 登录成功时，保存 token 到 sessionStorage
                    window.sessionStorage.setItem('token', res.token);

                    // 跳转到 'admin' 路由
                    this.$router.push('/admin');
                }
            } catch (error) {
                this.$message.error('登录请求失败，请重试');
            }
        }

    }
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
    padding: 20px 20px;
    box-sizing: border-box;
}

.loginBtn {
    display: flex;
    justify-content: flex-end;
}
</style>