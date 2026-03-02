import { useEffect, useState, useContext } from "react";

import GlobalContext from "../../context/global";

import Toolbar from "./toolbar";

const Editor = () => {

  const { state, dispatch } = useContext(GlobalContext.Context);
  const currentFile = state.currentOpenFile;
  const content = state.currentOpenFileContent;
  const setToc = (toc) => dispatch({ type: GlobalContext.Actions.SET_TOC, payload: toc });
  const setContent = (content) => dispatch({ type: GlobalContext.Actions.SET_CURRENT_OPEN_FILE_CONTENT, payload: content });

  const [isDirty, setIsDirty] = useState(false);

  useEffect(() => {
    if (currentFile) {
      fetch(`/api/ui/files/${currentFile}`)
        .then((res) => res.text())
        .then((data) => {
          const toc = generateToc(data);
          setToc(toc);
          setContent(data);
        })
        .catch((err) => console.error("Error fetching file content:", err));
    }
  }, [currentFile]);

  useEffect(() => {
    if (currentFile && content) {
      const toc = generateToc(content);
      setToc(toc);
    }
  }, [content]);


  return (
    <div className="container-fluid w-100 h-100 d-flex flex-column">
      <div className="row">
        <h2>{currentFile}</h2>
      </div>
      <div className="row">
        <Toolbar
          content={content}
          isDirty={isDirty}
          setIsDirty={setIsDirty}
        />
      </div>
      <div className="row flex-grow-1">
        <textarea
          className="editor border rounded p-2 w-100 h-100"
          // style={{ overflowY: 'scroll', whiteSpace: 'pre-wrap' }}
          // onSelect={(e) => {
          //   const selection = document.getSelection()
          //   setSelection(selection);
          // }}
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