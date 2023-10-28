<template>
    <div>
        <a-card>
            <h3>{{ id ? '编辑文章' : '新增文章' }}</h3>
            <a-form :model="artInfo" ref="artInfoRef" :rules="artInfoRules" :hideRequiredMark="true">
                <a-form-item label="文章标题" name="title">
                    <a-input v-model:value="artInfo.title" placeholder="请输入文章标题" style="width: 300px;"></a-input>
                </a-form-item>
                <a-form-item label="文章分类" name="cid">
                    <a-select v-model:value="artInfo.cid" placeholder="请选择分类" style="width: 300px;" @change="cateChange">
                        <a-select-option v-for="item in Catelist" :key="item.id" :value="item.id">{{ item.name
                        }}</a-select-option>
                    </a-select>
                </a-form-item>
                <a-form-item label="文章描述" name="desc">
                    <a-textarea v-model:value="artInfo.desc" placeholder="请输入文章描述" style="width: 300px;"></a-textarea>
                </a-form-item>
                <a-form-item label="文章缩略图" name="img">
                    <a-upload :defaultFileList="fileList" name="file" :action="upUrl" :headers="headers" @change="upChange"
                        listType="picture">
                        <a-button>
                            <UploadOutlined /> 上传图片
                        </a-button>
                        <br />
                        <template v-if="id">
                            <img :src="artInfo.img" style="width: 120px;height: 100px;">
                        </template>
                    </a-upload>
                </a-form-item>
                <a-form-item name="content">
                    <Editor v-model:value="artInfo.content" @tinymceinput="gettinymceinput"></Editor>
                </a-form-item>
                <a-form-item>
                    <a-button type="primary" style="margin-right: 15px;" @click="artOk(artInfo.id)">提交</a-button>
                    <a-button type="primary" danger ghost @click="addCancle">取消</a-button>
                </a-form-item>
            </a-form>
        </a-card>

    </div>
</template>


<script>
import { UploadOutlined } from '@ant-design/icons-vue';
import Editor from '../editor/EditorIndex.vue'
export default {
    props: ['id'],
    data() {
        return {
            artInfo: {
                id: 0,
                title: '',
                cid: undefined,
                desc: '',
                content: '',
                img: '',
            },
            Catelist: [],
            upUrl: "http://localhost:3000/api/v1/upload",
            headers: {},
            fileList: [],
            artInfoRules: {
                title: [{ required: true, message: '请输入文章标题', trigger: 'blur' }],
                cid: [{ required: true, message: '请选择文章分类', trigger: 'change' }],
                desc: [{ required: true, message: '请输入文章描述', trigger: 'blur' },
                { max: 120, message: "文章描述不能超过120个字符", trigger: "change" }],
                content: [{ required: true, message: '请输入文章内容', trigger: 'blur' }],
            }
        }
    },
    components: {
        UploadOutlined,
        Editor
    },
    created() {
        this.getCateList()
        this.headers = { Authorization: `Bearer ${window.sessionStorage.getItem('token')}` }
        if (this.id) {
            this.getArtInfo(this.id)
        }
    },
    methods: {
        // 查询文章信息
        async getArtInfo(id) {
            const { data: res } = await this.$http.get(`article/info/${id}`)
            if (res.status !== 200) return this.$message.error(res.message)
            this.artInfo = res.data
            this.artInfo.id = res.data.ID
        },

        // 获取分类列表
        async getCateList() {
            const { data: res } = await this.$http.get('category')
            if (res.status !== 200) return this.$message.error(res.message)
            this.Catelist = res.data
        },

        // 选择分类
        cateChange(value) {
            this.artInfo.cid = value
        },

        // 上传图片
        upChange(info) {
            if (info.file.status === 'done') {
                this.$message.success("图片上传成功");
                const imgUrl = info.file.response.url
                this.artInfo.img = imgUrl
            } else if (info.file.status === 'error') {
                this.$message.error("图片上传失败");
            }
        },

        // 提交文章 
        async artOk(id) {
            try {
                const valid = await this.$refs.artInfoRef.validate()
                if (!valid) return this.$message.error('输入非法数据，请重新输入')
                if (id === 0) {
                    const { data: res } = await this.$http.post('article/add', this.artInfo)
                    if (res.status !== 200) return this.$message.error(res.message)

                    this.$router.push('/artlist')
                    this.$message.success("文章添加成功")
                } else {
                    const { data: res } = await this.$http.put(`article/${id}`, this.artInfo)
                    if (res.status !== 200) return this.$message.error(res.message)

                    this.$router.push('/artlist')
                    this.$message.success("文章修改成功")
                }
            } catch (error) {
                this.$message.error('请求失败，请重试')
            }
        },
        addCancle() {
            this.$refs.artInfoRef.resetFields()
        },

        gettinymceinput(content) {
            this.artInfo.content = content
        }
    }
}
</script>