import React from 'react';

import Sidebar from './sidebar';
import Editor from './editor';

const MainContent = () => {
  const [currentFile, setCurrentFile] = React.useState(null);
  const [toc, setToc] = React.useState([]);

  return (
    <main className="container-fluid">
      <div className="row">
        <div className="col-2">
            <Sidebar
              currentFile={currentFile}
              setCurrentFile={setCurrentFile}
              toc={toc}
            />
        </div>
        { currentFile && <div className="col-10">
            <Editor currentFile={currentFile} setCurrentFile={setCurrentFile} setToc={setToc} />
        </div>}
      </div>
    </main>
  );
};

export default MainContent;