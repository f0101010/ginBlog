<template>
    <div>
        <a-card>
            <a-row :gutter="20">
                <a-col :span="6">
                    <a-input-search v-model:value="queryParam.title" placeholder="输入文章标题查找" enter-button
                        @search="getArtList" />
                </a-col>
                <a-col :span="4">
                    <a-button type="primary" @click="addArtVisible = true">新增文章</a-button>
                </a-col>

                <a-col :span="6" :offset="4">
                    <a-select defaultValue="请选择分类" style="width: 200px;" @change="cateChange">
                        <a-select-option v-for="item in Catelist" :key="item.id" :value="item.id">
                            {{ item.name }}
                        </a-select-option>
                    </a-select>
                </a-col>
            </a-row>
            <a-table rowKey="Artname" :columns="columns" :pagination="pagination" :dataSource="Artlist" bordered allowClear
                @change="handleTableChange">
                <template v-slot:bodyCell="{ column, record }">
                    <template v-if="column.dataIndex === 'img'">
                        <span class="artImg">
                            <img :src="record.img" alt="">
                        </span>
                    </template>
                    <template v-if="column.dataIndex === 'action'">
                        <div class="actionSlot">
                            <a-button type="primary" style="margin-right:16px;" @click="editArt(record.ID)">编辑</a-button>
                            <a-button type="primary" danger ghost style="margin-right:16px;"
                                @click="deleteArt(record.ID)">删除</a-button>
                        </div>
                    </template>
                </template>
            </a-table>
        </a-card>
    </div>
</template>


<script>

const columns = [
    {
        title: 'ID',
        dataIndex: 'ID',
        width: '5%',
        key: 'id',
        align: 'center',
    },
    {
        title: '分类',
        dataIndex: ['Category', 'name'],
        width: '10%',
        key: 'name',
        align: 'center',
    },
    {
        title: '文章标题',
        dataIndex: 'title',
        width: '10%',
        key: 'title',
        align: 'center',
    },
    {
        title: '文章描述',
        dataIndex: 'desc',
        width: '20%',
        key: 'desc',
        align: 'center',
    },
    {
        title: '缩略图',
        dataIndex: 'img',
        width: '20%',
        key: 'img',
        align: 'center',
    },
    {
        title: '操作',
        width: '20%',
        key: 'action',
        dataIndex: 'action',
        align: 'center',
    },

]

export default {
    data() {
        return {
            pagination: {
                pageSizeOptions: ['5', '10', '20'],
                pageSize: 5,
                total: 0,
                showSizeChanger: true,
                showTotal: (total) => `共${total}条`,
            },
            Artlist: [],
            Catelist: [],
            columns,
            queryParam: {
                title: '',
                pagenum: 1,
                pagesize: 5,
            },
            layout: {
                labelCol: {
                    span: 4,
                },
            },
            AllowClear: false,
        }
    },
    created() {
        this.getArtList()
        this.getCateList()
    },
    methods: {
        // 获取文章列表
        async getArtList() {
            const { data: res } = await this.$http.get('article', {
                params: {
                    title: this.queryParam.title,
                    pagesize: this.queryParam.pagesize,
                    pagenum: this.queryParam.pagenum,
                },
            })
            if (res.status !== 200) return this.$message.error(res.message)
            this.Artlist = res.data
            this.pagination.total = res.total
        },

        // 获取分类
        async getCateList() {
            const { data: res } = await this.$http.get('category')
            if (res.status !== 200) return this.$message.error(res.message)
            this.Catelist = res.data
            this.pagination.total = res.total
        },
        // 更改页码
        handleTableChange(pagination) {
            var pager = { ...this.pagination }
            pager.current = pagination.current
            pager.pageSize = pagination.pageSize
            this.queryParam.pagesize = pagination.pageSize
            this.queryParam.pagenum = pagination.current


            if (pagination.pageSize !== this.pagination.pageSize) {
                this.queryParam.pagenum = 1
                pager.current = 1
            }
            this.pagination = pager
            this.getArtList()
        },

        // 删除文章
        deleteArt(id) {
            this.$confirm({
                title: '确定要删除该文章吗?',
                content: '一旦删除，无法恢复',
                okText: 'Yes',
                okType: 'danger',
                cancelText: 'No',
                onOk: async () => {
                    const res = await this.$http.delete(`article/${id}`)
                    if (res.status !== 200) return this.$message.error(res.message)
                    this.$message.success('删除成功')
                    this.getArtList()
                },
                onCancel: () => {
                    this.$message.info('已取消删除')
                },
            });

        },

        // 查询分类下的文章
        cateChange(value) {
            this.getCateArt(value)
        },
        async getCateArt(id) {
            const { data: res } = await this.$http.get(`article/list/${id}`)
            if (res.status !== 200) return this.$message.error(res.message)
            this.Artlist = res.data
            this.pagination.total = res.total
        }
    },
}
</script>


<style scoped>
.actionSlot {
    display: flex;
    justify-content: center;
}

.artImg {
    height: 100%;
    width: 100%;
}

.artImg img {
    width: 100px;
    height: 80px;
}
</style>