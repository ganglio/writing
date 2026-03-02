import { useEffect, useContext } from "react";

import GlobalContext from "../context/global";

const Files = () => {
  const { state, dispatch } = useContext(GlobalContext.Context);

  useEffect(() => {
    fetch("http://localhost:3000/api/ui/files")
      .then((res) => res.json())
      .then((data) => dispatch({ type: GlobalContext.Actions.SET_FILE_LIST, payload: data }))
      .catch((err) => console.error("Error fetching files:", err));
  }, []);

  return (
    <section className="overflow-scroll">
      <span className="fw-bold">File Explorer</span>
      <ul className="list-group list-group-flush">
        {state.files.map((file, index) => (
          <li key={index} className={`list-group-item ${file === state.currentOpenFile ? 'active' : ''}`}>
            <button
              className="nav-link"
              onClick={() => dispatch({ type: GlobalContext.Actions.SET_CURRENT_OPEN_FILE, payload: file })}>
              {file}
            </button>
          </li>
        ))}
      </ul>
    </section>
  );
};

export default Files;