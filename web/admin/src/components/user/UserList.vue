<template>
    <div>
        <h3>用户列表页面</h3>
        <a-card>
            <a-row :gutter="20">
                <a-col :span="6">
                    <a-input-search placeholder="输入用户名查找" enter-button />
                </a-col>
                <a-col :span="4">
                    <a-button type="primary">新增用户</a-button>
                </a-col>
            </a-row>
            <a-table rowKey="username" :columns="columns" :pagination="paginationOption" :dataSource="userlist" bordered>
                <template #name="{ }">
                    <div class="actionSlot">
                        <a-button type="primary" style="margin-right: 16px;">编辑</a-button>
                        <a-button type="primary" danger ghost>删除</a-button>
                    </div>
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
        width: '10%',
        key: 'id',
        align: 'center',
    },
    {
        title: '用户名',
        dataIndex: 'username',
        width: '20%',
        key: 'username',
        align: 'center',
    },
    {
        title: '角色',
        dataIndex: 'role',
        width: '20%',
        key: 'role',
        align: 'center',
        customRender: (text) => {
            return text.value == 1 ? '管理员' : '普通用户'
        }
    },
    {
        title: '操作',
        width: '30%',
        key: 'action',
        slots: { customRender: 'name' }
    },
]

export default {
    data() {
        return {
            paginationOption: {
                pageSizeOptions: ['5', '10', '20'],
                defaultCurrent: 1,
                defaultPageSize: 5,
                total: 0,
                showSizeChanger: true,
                showTotal: (total) => `共${total}条`,
                onChange: (current, pageSize) => {
                    this.paginationOption.defaultCurrent = current
                    this.paginationOption.defaultPageSize = pageSize
                    this.getUserList()
                },
                showSizeChange: (current, size) => {
                    this.paginationOption.defaultCurrent = current
                    this.paginationOption.defaultPageSize = size
                    this.getUserList()
                },
            },
            userlist: [],
            columns,
        }
    },
    created() {
        this.getUserList()
    },
    methods: {
        async getUserList() {
            const { data: res } = await this.$http.get('users', {
                params: {
                    pagesize: this.PageSize,
                    pagenum: this.defaultCurrent,
                },
            })
            if (res.status !== 200) return this.$message.error(res.message)
            this.userlist = res.data
            this.paginationOption.total = res.total
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