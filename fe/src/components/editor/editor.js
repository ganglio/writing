import { useEffect, useState } from "react";
import Toolbar from "./toolbar";

const Editor = ({ currentFile, setToc }) => {

  const [content, setContent] = useState('');
  const [selection, setSelection] = useState(null);
  const [isDirty, setIsDirty] = useState(false);

  useEffect(() => {
    if (currentFile) {
      fetch(`/api/files/${currentFile.filename}`)
        .then((res) => res.text())
        .then((data) => {
          const toc = generateToc(data);
          setToc(toc);
          setContent(data)
        })
        .catch((err) => console.error("Error fetching file content:", err));
    }
  }, [currentFile]);

  useEffect(() => {
    if (currentFile) {
      const toc = generateToc(content);
      setToc(toc);
    }
  }, [content]);


  return (
    <div className="container-fluid w-100 h-100 d-flex flex-column">
      <div className="row">
        <h2>{currentFile.filename}</h2>
      </div>
      <div className="row">
        <Toolbar
          selection={selection}
          content={content}
          setContent={setContent}
          isDirty={isDirty}
          setIsDirty={setIsDirty}
        />
      </div>
      <div className="row flex-grow-1">
        <textarea
          className="editor border rounded p-2 w-100 h-100"
          // style={{ overflowY: 'scroll', whiteSpace: 'pre-wrap' }}
          onSelect={(e) => {
            const selection = document.getSelection()
            setSelection(selection);
          }}
          onInput={(e) => {
            setIsDirty(true);
          }}
          value={content}
        />
      </div>
    </div>
  );
};

const generateToc = (data) => {
  const toc = data
    .split("\n")
    .filter(a=>a)
    .filter(s=>/^#/.test(s))
    .map(s=>[...s.matchAll(/(#{1,}) (.*)/g)])
    .map(m=>({title: m[0][2], level: m[0][1].length }));
    
  return toc;
}

export default Editor;