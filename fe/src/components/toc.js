import { useState, useEffect } from "react";

const ToC = ({ toc }) => {
  console.log(toc)
  return (
    toc && <aside className="mt-4">
      <span className="fw-bold">Table of Contents</span>
      <ul className="list-group list-group-flush">
        {toc.map((section, index) => (
          <li key={index} className={`list-group-item`}>
            <a className="nav-link" href={`#${section.title}`}>{section.title}</a>
          </li>
        ))}
      </ul>
    </aside>
  );
};

export default ToC;