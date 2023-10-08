<template>
    <div>
        <a-card>
            <a-row :gutter="20">
                <a-col :span="6">
                    <a-input-search v-model:value="queryParam.username" placeholder="输入用户名查找" enter-button
                        @search="getUserList" />
                </a-col>
                <a-col :span="4">
                    <a-button type="primary" @click="addUserVisible = true">新增用户</a-button>
                </a-col>
            </a-row>
            <a-table rowKey="username" :columns="columns" :pagination="pagination" :dataSource="userlist" bordered
                allowClear @change="handleTableChange">
                <template v-slot:bodyCell="{ column, record }">
                    <template v-if="column.dataIndex === 'action'">
                        <div class="actionSlot">
                            <a-button type="primary" style="margin-right:16px;" @click="editUser(record.ID)">编辑</a-button>
                            <a-button type="primary" danger ghost style="margin-right:16px;"
                                @click="deleteUser(record.ID)">删除</a-button>
                            <a-button @click="resetPwd(record.ID)">重置密码</a-button>
                        </div>
                    </template>
                    <template v-if="column.dataIndex === 'role'">
                        <span>{{ record.role === 1 ? '管理员' : "普通用户" }}</span>
                    </template>
                </template>
            </a-table>
        </a-card>
        <!-- 新增用户 -->
        <a-modal closable title="新增用户" :open="addUserVisible" @ok="addUserOk" @cancel="addUserCancle" destroyOnClose
            width="50%">
            <a-form :model="newUser" :rules="addUserRules" ref="addUserRef" v-bind="layout">
                <a-form-item has-feedback label="用户名：" name="username">
                    <a-input v-model:value="newUser.username"></a-input>
                </a-form-item>
                <a-form-item has-feedback label="密码：" name="password">
                    <a-input-password v-model:value="newUser.password"></a-input-password>
                </a-form-item>
                <a-form-item has-feedback label="确认密码：" name="confirmPwd">
                    <a-input-password v-model:value="newUser.confirmPwd"></a-input-password>
                </a-form-item>
            </a-form>
        </a-modal>

        <!-- 编辑用户 -->
        <a-modal closable title="编辑用户" :open="editUserVisible" @ok="editUserOk" @cancel="editUserCancle" width="50%">
            <a-form :model="userInfo" :rules="userRules" ref="editUserRef" v-bind="layout">
                <a-form-item has-feedback label="用户名：" name="username">
                    <a-input v-model:value="userInfo.username"></a-input>
                </a-form-item>
                <a-form-item label="是否为管理员：">
                    <a-switch :checked="IsAdmin" checked-children="是" un-checked-children="否" @change="adminChange" />
                </a-form-item>
            </a-form>
        </a-modal>

        <!-- 重置密码 -->
        <a-modal closable title="重置密码" :open="resetPwdVisible" @ok="resetPwdOk" @cancel="resetPwdCancle" width="50%">
            <a-form :model="newPwd" :rules="resetPwdRules" ref="resetPwdRef" v-bind="layout">
                <a-form-item has-feedback label="用户名：" name="username">
                    <div>{{ userInfo.username }}</div>
                </a-form-item>
                <a-form-item has-feedback label="密码：" name="password">
                    <a-input-password v-model:value="newPwd.password"></a-input-password>
                </a-form-item>
                <a-form-item has-feedback label="确认密码：" name="confirmPwd">
                    <a-input-password v-model:value="newPwd.confirmPwd"></a-input-password>
                </a-form-item>
            </a-form>
        </a-modal>
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
    },
    {
        title: '操作',
        width: '30%',
        key: 'action',
        dataIndex: 'action',
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
            userlist: [],
            columns,
            queryParam: {
                username: '',
                pagenum: 1,
                pagesize: 5,
            },
            userInfo: {
                id: 0,
                username: '',
                password: '',
                confirmPwd: '',
                role: 2,
            },
            newUser: {
                id: 0,
                username: '',
                password: '',
                confirmPwd: '',
                role: 2,
            },
            newPwd: {
                password: '',
                confirmPwd: '',
            },
            addUserVisible: false,
            userRules: {
                username: [{
                    validator: async (_rule, value) => {
                        if (value === '') {
                            return Promise.reject("请输入用户名");
                        } else if (value.length < 4 || value.length > 12) {
                            return Promise.reject("用户名长度为4-12位");
                        } else {
                            return Promise.resolve();
                        }
                    },
                    trigger: 'blur',
                }],
                password: [
                    {
                        validator: async (_rule, value) => {
                            if (value === '') {
                                return Promise.reject("请输入密码");
                            } else if (value.length < 6 || value.length > 20) {
                                return Promise.reject("密码长度为6-20位");
                            } else {
                                return Promise.resolve();
                            }
                        },
                        trigger: 'blur',
                    }
                ],
                confirmPwd: [{
                    validator: async (_rule, value) => {
                        if (value === '') {
                            return Promise.reject("请输入确认密码");
                        } else if (value !== this.userInfo.password) {
                            return Promise.reject("两次密码不一致");
                        } else {
                            return Promise.resolve();
                        }
                    },
                    trigger: 'blur',
                }]
            },
            addUserRules: {
                username: [{
                    validator: async (_rule, value) => {
                        if (value === '') {
                            return Promise.reject("请输入用户名");
                        } else if (value.length < 4 || value.length > 12) {
                            return Promise.reject("用户名长度为4-12位");
                        } else {
                            return Promise.resolve();
                        }
                    },
                    trigger: 'blur',
                }],
                password: [
                    {
                        validator: async (_rule, value) => {
                            if (value === '') {
                                return Promise.reject("请输入密码");
                            } else if (value.length < 6 || value.length > 20) {
                                return Promise.reject("密码长度为6-20位");
                            } else {
                                return Promise.resolve();
                            }
                        },
                        trigger: 'blur',
                    }
                ],
                confirmPwd: [{
                    validator: async (_rule, value) => {
                        if (value === '') {
                            return Promise.reject("请输入确认密码");
                        } else if (value !== this.newUser.password) {
                            return Promise.reject("两次密码不一致");
                        } else {
                            return Promise.resolve();
                        }
                    },
                    trigger: 'blur',
                }]
            },
            resetPwdRules: {
                password: [
                    {
                        validator: async (_rule, value) => {
                            if (value === '') {
                                return Promise.reject("请输入密码");
                            } else if (value.length < 6 || value.length > 20) {
                                return Promise.reject("密码长度为6-20位");
                            } else {
                                return Promise.resolve();
                            }
                        },
                        trigger: 'blur',
                    }
                ],
                confirmPwd: [{
                    validator: async (_rule, value) => {
                        if (value === '') {
                            return Promise.reject("请输入确认密码");
                        } else if (value !== this.resetPwd.password) {
                            return Promise.reject("两次密码不一致");
                        } else {
                            return Promise.resolve();
                        }
                    },
                    trigger: 'blur',
                }]
            },
            layout: {
                labelCol: {
                    span: 4,
                },
            },
            editUserVisible: false,
            resetPwdVisible: false,
        }
    },
    created() {
        this.getUserList()
    },
    computed: {
        IsAdmin: function () {
            if (this.userInfo.role === 1) {
                return true
            } else {
                return false
            }
        },
    },
    methods: {
        async getUserList() {
            const { data: res } = await this.$http.get('users', {
                params: {
                    username: this.queryParam.username,
                    pagesize: this.queryParam.pagesize,
                    pagenum: this.queryParam.pagenum,
                },
            })
            if (res.status !== 200) return this.$message.error(res.message)
            this.userlist = res.data
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
            this.getUserList()
        },

        // 删除用户
        deleteUser(id) {
            this.$confirm({
                title: '确定要删除该用户吗?',
                content: '一旦删除，无法恢复',
                okText: 'Yes',
                okType: 'danger',
                cancelText: 'No',
                onOk: async () => {
                    const res = await this.$http.delete(`user/${id}`)
                    if (res.status !== 200) return this.$message.error(res.message)
                    this.$message.success('删除成功')
                    this.getUserList()
                },
                onCancel: () => {
                    this.$message.info('已取消删除')
                },
            });

        },
        // 新增用户 
        async addUserOk() {
            try {
                const valid = await this.$refs.addUserRef.validate()
                if (!valid) {
                    this.$message.error('输入非法数据，请重新输入')
                    return;
                }
                const { data: res } = await this.$http.post('user/add', {
                    username: this.newUser.username,
                    password: this.newUser.password,
                    role: this.newUser.role
                })
                if (res.status !== 200) {
                    return this.$message.error(res.message)
                } else {
                    this.addUserVisible = false
                    this.$message.success('新增用户成功')
                    this.getUserList()
                }
            } catch (error) {
                this.$message.error('新增用户失败，请重试')
            }
        },
        addUserCancle() {
            this.$refs.addUserRef.resetFields()
            this.addUserVisible = false
            this.$message.info('已取消新增')
        },
        adminChange(checked) {
            if (checked) {
                this.userInfo.role = 1
            } else {
                this.userInfo.role = 2
            }
        },

        // 编辑用户
        async editUser(id) {
            this.editUserVisible = true
            const { data: res } = await this.$http.get(`user/${id}`)
            this.userInfo = res.data
            this.userInfo.id = id
        },
        async editUserOk() {
            try {
                const valid = await this.$refs.editUserRef.validate()
                if (!valid) {
                    this.$message.error('输入非法数据，请重新输入')
                    return;
                }
                const { data: res } = await this.$http.put(`user/${this.userInfo.id}`, {
                    username: this.userInfo.username,
                    role: this.userInfo.role
                })
                if (res.status !== 200) {
                    return this.$message.error(res.message)
                } else {
                    this.editUserVisible = false
                    this.$message.success('更新用户信息成功')
                    this.getUserList()
                }
            } catch (error) {
                this.$message.error('更新用户信息失败，请重试')
            }
        },
        editUserCancle() {
            this.$refs.editUserRef.resetFields()
            this.editUserVisible = false
            this.$message.info('已取消编辑')
        },

        // 重置密码
        async resetPwd(id) {
            this.resetPwdVisible = true
            const { data: res } = await this.$http.get(`user/${id}`)
            this.userInfo = res.data
            this.userInfo.id = id
        },
        async resetPwdOk() {
            try {
                const valid = await this.$refs.resetPwdRef.validate()
                if (!valid) {
                    this.$message.error('输入非法数据，请重新输入')
                    return;
                }
                const { data: res } = await this.$http.put(`reset/${this.userInfo.id}`, {
                    password: this.newPwd.password,
                })
                if (res.status !== 200) {
                    return this.$message.error(res.message)
                } else {
                    this.resetPwdVisible = false
                    this.$message.success('重置密码成功')
                    this.getUserList()
                }
            } catch (error) {
                this.$message.error('重置密码失败，请重试')
            }
        },
        resetPwdCancle() {
            this.$refs.resetPwdRef.resetFields()
            this.resetPwdVisible = false
            this.$message.info('已取消重置密码')
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