import React from 'react';

import Sidebar from './sidebar';
import Editor from './editor/editor';

const MainContent = () => {
  const [currentFile, setCurrentFile] = React.useState(null);
  const [toc, setToc] = React.useState([]);

  return (
    <main className="container-fluid d-flex">
      <div className="col-2">
          <Sidebar
            currentFile={currentFile}
            setCurrentFile={setCurrentFile}
            toc={toc}
          />
      </div>
      { currentFile && <div className="col flex-grow-1">
          <Editor currentFile={currentFile} setCurrentFile={setCurrentFile} setToc={setToc} />
      </div>}
    </main>
  );
};

export default MainContent;