<template>
    <div>
        <a-card>
            <a-row :gutter="20">
                <a-col :span="4">
                    <a-button type="primary" @click="addCateVisible = true">新增分类</a-button>
                </a-col>
            </a-row>
            <a-table rowKey="Catename" :columns="columns" :pagination="pagination" :dataSource="Catelist" bordered
                allowClear @change="handleTableChange">
                <template v-slot:bodyCell="{ column, record }">
                    <template v-if="column.dataIndex === 'action'">
                        <div class="actionSlot">
                            <a-button type="primary" style="margin-right:16px;" @click="editCate(record.id)">编辑</a-button>
                            <a-button type="primary" danger ghost style="margin-right:16px;"
                                @click="deleteCate(record.id)">删除</a-button>
                        </div>
                    </template>
                </template>
            </a-table>
        </a-card>
        <!-- 新增分类 -->
        <a-modal closable title="新增分类" :open="addCateVisible" @ok="addCateOk" @cancel="addCateCancle" destroyOnClose
            width="50%">
            <a-form :model="newCate" :rules="addCateRules" ref="addCateRef" v-bind="layout">
                <a-form-item has-feedback label="分类名称：" name="Catename">
                    <a-input v-model:value="newCate.name"></a-input>
                </a-form-item>
            </a-form>
        </a-modal>

        <!-- 编辑分类 -->
        <a-modal closable title="编辑分类" :open="editCateVisible" @ok="editCateOk" @cancel="editCateCancle" width="50%">
            <a-form :model="CateInfo" :rules="CateRules" ref="editCateRef" v-bind="layout">
                <a-form-item has-feedback label="分类名称：" name="Catename">
                    <a-input v-model:value="CateInfo.name"></a-input>
                </a-form-item>
            </a-form>
        </a-modal>
    </div>
</template>


<script>

const columns = [
    {
        title: 'ID',
        dataIndex: 'id',
        width: '10%',
        key: 'id',
        align: 'center',
    },
    {
        title: '分类名',
        dataIndex: 'name',
        width: '20%',
        key: 'name',
        align: 'center',
    },
    {
        title: '操作',
        width: '30%',
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
            Catelist: [],
            columns,
            queryParam: {
                pagenum: 1,
                pagesize: 5,
            },
            CateInfo: {
                id: 0,
                name: '',
            },
            newCate: {
                name: '',
            },
            addCateVisible: false,
            CateRules: {
                name: [{
                    validator: async (_rule, value) => {
                        if (value === '') {
                            return Promise.reject("请输入分类名");
                        } else {
                            return Promise.resolve();
                        }
                    },
                    trigger: 'blur',
                }],

            },
            addCateRules: {
                name: [{
                    validator: async (_rule, value) => {
                        if (value === '') {
                            return Promise.reject("请输入分类名");
                        } else {
                            return Promise.resolve();
                        }
                    },
                    trigger: 'blur',
                }],
            },
            layout: {
                labelCol: {
                    span: 4,
                },
            },
            editCateVisible: false,
            resetPwdVisible: false,
        }
    },
    created() {
        this.getCateList()
    },
    methods: {
        // 获取分类列表
        async getCateList() {
            const { data: res } = await this.$http.get('category', {
                params: {
                    pagesize: this.queryParam.pagesize,
                    pagenum: this.queryParam.pagenum,
                },
            })
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
            this.getCateList()
        },

        // 删除分类 
        deleteCate(id) {
            this.$confirm({
                title: '确定要删除该分类吗?',
                content: '一旦删除，无法恢复',
                okText: 'Yes',
                okType: 'danger',
                cancelText: 'No',
                onOk: async () => {
                    const res = await this.$http.delete(`category/${id}`)
                    if (res.status !== 200) return this.$message.error(res.message)
                    this.$message.success('删除成功')
                    this.getCateList()
                },
                onCancel: () => {
                    this.$message.info('已取消删除')
                },
            });

        },
        // 新增分类 
        async addCateOk() {
            try {
                const valid = await this.$refs.addCateRef.validate()
                if (!valid) {
                    this.$message.error('输入非法数据，请重新输入')
                    return;
                }
                const { data: res } = await this.$http.post('category/add', {
                    name: this.newCate.name,
                })
                if (res.status !== 200) {
                    return this.$message.error(res.message)
                } else {
                    this.addCateVisible = false
                    this.$message.success('添加分类成功')
                    await this.getCateList()
                }
            } catch (error) {
                this.$message.error('添加分类失败，请重试')
            }
        },
        addCateCancle() {
            this.$refs.addCateRef.resetFields()
            this.addCateVisible = false
            this.$message.info('已取消新增')
        },

        // 编辑分类
        async editCate(id) {
            this.editCateVisible = true
            const { data: res } = await this.$http.get(`category/${id}`)
            this.CateInfo = res.data
            this.CateInfo.id = id
        },
        async editCateOk() {
            try {
                const valid = await this.$refs.editCateRef.validate()
                if (!valid) {
                    this.$message.error('输入非法数据，请重新输入')
                    return;
                }
                const { data: res } = await this.$http.put(`category/${this.CateInfo.id}`, {
                    name: this.CateInfo.name,
                })
                if (res.status !== 200) {
                    return this.$message.error(res.message)
                } else {
                    this.editCateVisible = false
                    this.$message.success('更新分类成功')
                    this.getCateList()
                }
            } catch (error) {
                this.$message.error('更新分类失败，请重试')
            }
        },
        editCateCancle() {
            this.$refs.editCateRef.resetFields()
            this.editCateVisible = false
            this.$message.info('已取消编辑')
        },

    },
}
</script>


<style scoped>
.actionSlot {
    display: flex;
    justify-content: center;
}
</style>