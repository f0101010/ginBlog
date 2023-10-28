<template>
    <div>
        <Editor id="tinymce" :init="init" v-model="content"></Editor>
    </div>
</template>

<script>
import Editor from '@tinymce/tinymce-vue'
import 'tinymce/tinymce'
import 'tinymce/icons/default/icons'
import 'tinymce/themes/silver'
import 'tinymce/models/dom/model'

// 注册插件
import 'tinymce/plugins/preview/plugin'
import 'tinymce/plugins/wordcount/plugin'
import 'tinymce/plugins/code/plugin'
import 'tinymce/plugins/image/plugin'


export default {
    components: {
        Editor
    },
    props: {
        value: {
            type: String,
            default: ''
        }
    },
    data() {
        return {
            init: {
                language: 'zh-Hans',
                height: '500px',
                margin: '0',
                pading: '0',
                plugins: 'preview wordcount code image',
                toolbar: 'undo redo | formatselect | bold italic backcolor | code | link image |',
                branding: false,
                skin_url: 'admin/tinymce/skins/ui/oxide',
                content_css: 'admin/tinymce/skins/content/default/content.min.css',

                images_upload_handler: (blobInfo) => {
                    return new Promise((resolve, reject) => {
                        let formdata = new FormData();
                        formdata.append('file', blobInfo.blob(), blobInfo.filename());

                        this.$http.post('upload', formdata)
                            .then(response => {
                                const res = response.data;
                                if (res.status !== 200) {
                                    reject(res.message);
                                } else {
                                    resolve(res.url);
                                }
                            })
                            .catch(error => {
                                reject(error);
                            });
                    });
                },
            },
            content: this.value,
        }
    },
    watch: {
        value(newV) {
            this.content = newV
        },
        content(newV) {
            this.$emit('tinymceinput', newV)
        },
    },

}
</script>

<style></style>