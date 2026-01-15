import { useEffect, useState } from "react";

const Editor = ({ currentFile, setToc }) => {

  const [content, setContent] = useState('');
  useEffect(() => {
    if (currentFile) {
      fetch(`http://localhost:3000/api/files/${currentFile.filename}`)
        .then((res) => res.text())
        .then((data) => {
          const toc = generateToc(data);
          setToc(toc);
          setContent(data)
        })
        .catch((err) => console.error("Error fetching file content:", err));
    }
  }, [currentFile]);

  return (
    <div className="container-fluid w-100 mh-100">
      <h2>{currentFile.filename}</h2>
      <div className="editor border rounded p-3 h-100" style={{ overflowY: 'scroll', whiteSpace: 'pre-wrap' }}>
        {content}
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