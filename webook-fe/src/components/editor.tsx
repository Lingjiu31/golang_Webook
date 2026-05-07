import '@wangeditor/editor/dist/css/style.css'
import React, { useState, useEffect } from 'react'
import { Editor, Toolbar } from '@wangeditor/editor-for-react'
import { IDomEditor, IEditorConfig, IToolbarConfig } from '@wangeditor/editor'

interface Props {
    html: any,
    setHtmlFn: any
}

function WangEditor(props: Props) {
    const [editor, setEditor] = useState<IDomEditor | null>(null)

    const {html, setHtmlFn} = props

    const toolbarConfig: Partial<IToolbarConfig> = { }

    const editorConfig: Partial<IEditorConfig> = {
        placeholder: '请输入内容...',
    }

    useEffect(() => {
        return () => {
            if (editor == null) return
            editor.destroy()
            setEditor(null)
        }
    }, [editor])

    return (
        <div style={{ border: '1px solid #ced0d4', zIndex: 100, borderRadius: '8px', overflow: 'hidden' }}>
            <Toolbar
                editor={editor}
                defaultConfig={toolbarConfig}
                mode="default"
                style={{ borderBottom: '1px solid #dee3e9', backgroundColor: '#f1f4f7' }}
            />
            <Editor
                defaultConfig={editorConfig}
                value={html}
                onCreated={setEditor}
                onChange={editor => setHtmlFn(editor.getHtml())}
                mode="default"
                style={{ height: '500px', overflowY: 'hidden', backgroundColor: '#ffffff' }}
            />
        </div>
    )
}
export default WangEditor
