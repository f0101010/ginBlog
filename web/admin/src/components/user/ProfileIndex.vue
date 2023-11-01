<template>
    <div>
        <a-card>
            <h3>个人设置</h3>
            <a-form v-bind="layout">
                <a-form-item label="作者名称">
                    <a-input v-model:value="profileInfo.name" style="width: 300px;"></a-input>
                </a-form-item>

                <a-form-item label="个人简介">
                    <a-textarea v-model:value="profileInfo.desc" style="width: 300px;"></a-textarea>
                </a-form-item>

                <a-form-item label="QQ号码">
                    <a-input v-model:value="profileInfo.qq_chat" style="width: 300px;"></a-input>
                </a-form-item>

                <a-form-item label="微信">
                    <a-input v-model:value="profileInfo.wechat" style="width: 300px;"></a-input>
                </a-form-item>

                <a-form-item label="音乐">
                    <a-input v-model:value="profileInfo.music" style="width: 300px;"></a-input>
                </a-form-item>

                <a-form-item label="哔哩哔哩">
                    <a-input v-model:value="profileInfo.bili" style="width: 300px;"></a-input>
                </a-form-item>

                <a-form-item label="Email">
                    <a-input v-model:value="profileInfo.email" style="width: 300px;"></a-input>
                </a-form-item>

                <a-form-item label="头像" name="img">
                    <a-upload :defaultFileList="fileList" name="file" :action="upUrl" :headers="headers"
                        @change="avatarChange" listType="picture">
                        <a-button>
                            <UploadOutlined /> 上传图片
                        </a-button>
                        <br />
                        <template v-if="profileInfo.avatar">
                            <img :src="profileInfo.avatar" style="width: 120px;height: 100px;">
                        </template>
                    </a-upload>
                </a-form-item>

                <a-form-item label="头像背景图" name="img">
                    <a-upload :defaultFileList="fileList" name="file" :action="upUrl" :headers="headers" @change="imgChange"
                        listType="picture">
                        <a-button>
                            <UploadOutlined /> 上传图片
                        </a-button>
                        <br />
                        <template v-if="profileInfo.img">
                            <img :src="profileInfo.img" style="width: 120px;height: 100px;">
                        </template>
                    </a-upload>
                </a-form-item>

                <a-form-item>
                    <a-button type="primary" style="margin-right: 15px;" @click="updateProfile">更新</a-button>
                </a-form-item>
            </a-form>
        </a-card>
    </div>
</template>



<script>
import { UploadOutlined } from '@ant-design/icons-vue';
export default {
    data() {
        return {
            profileInfo: {
                id: 1,
                name: '',
                desc: '',
                qq_chat: '',
                wechat: '',
                music: '',
                bili: '',
                email: '',
                img: '',
                avatar: '',
            },
            upUrl: "http://localhost:3000/api/v1/upload",
            headers: {},
            fileList: [],
            layout: {
                labelCol: {
                    span: 2,
                },
            },
        }
    },
    components: {
        UploadOutlined,
    },
    created() {
        this.getProfileInfo()
        this.headers = { Authorization: `Bearer ${window.sessionStorage.getItem('token')}` }
    },
    methods: {
        // 获取个人信息
        async getProfileInfo() {
            const { data: res } = await this.$http.get(`profile/${this.profileInfo.id}`)
            if (res.status !== 200) return this.$message.error(res.message)
            this.profileInfo = res.data
        },

        // 上传头像
        avatarChange(info) {
            if (info.file.status === 'done') {
                this.$message.success("图片上传成功");
                const imgUrl = info.file.response.url
                this.profileInfo.avatar = imgUrl
            } else if (info.file.status === 'error') {
                this.$message.error("图片上传失败");
            }
        },

        // 上传头像背景图
        imgChange(info) {
            if (info.file.status === 'done') {
                this.$message.success("图片上传成功");
                const imgUrl = info.file.response.url
                this.profileInfo.img = imgUrl
            } else if (info.file.status === 'error') {
                this.$message.error("图片上传失败");
            }
        },

        // 更新个人信息
        async updateProfile() {
            const { data: res } = await this.$http.put(`profile/${this.profileInfo.id}`, this.profileInfo)
            if (res.status !== 200) return this.$message.error(res.message)
            this.$message.success('更新成功')
            this.$router.push('/index')
        },
    }
}
</script>

<style scoped></style>
