import React from 'react';

import GlobalContext from '../context/global';

import Sidebar from './sidebar';
import Editor from './editor/editor';

const MainContent = () => {
  const { state, dispatch } = React.useContext(GlobalContext.Context);

  return (
    <main className="container-fluid d-flex">
      <div className="col-2">
          <Sidebar />
      </div>
      { state.currentOpenFile && <div className="col flex-grow-1">
          <Editor currentFile={state.currentOpenFile} setCurrentFile={(file) => dispatch({ type: GlobalContext.Actions.SET_CURRENT_OPEN_FILE, payload: file })} setToc={(toc) => dispatch({ type: GlobalContext.Actions.SET_TOC, payload: toc })  } />
      </div> }
    </main>
  );
};

export default MainContent;