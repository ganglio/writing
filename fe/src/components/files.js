import { useState, useEffect } from "react";

const Files = ({ currentFile, setCurrentFile }) => {
  const [files, setFiles] = useState([]);

  useEffect(() => {
    fetch("http://localhost:3000/api/files")
      .then((res) => res.json())
      .then((data) => setFiles(data))
      .catch((err) => console.error("Error fetching files:", err));
  }, []);

  return (
    <section className="overflow-scroll">
      <span className="fw-bold">File Explorer</span>
      <ul className="list-group list-group-flush">
        {files.map((file, index) => (
          <li key={index} className={`list-group-item ${file === currentFile ? 'active' : ''}`}>
            <button className="nav-link" onClick={() => setCurrentFile({...currentFile, filename: file})}>{file}</button>
          </li>
        ))}
      </ul>
    </section>
  );
};

export default Files;