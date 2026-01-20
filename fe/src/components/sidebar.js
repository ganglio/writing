import Files from "./files";
import ToC from "./toc";

const Sidebar = ({ currentFile, setCurrentFile, toc }) => {
  return (
    <aside className="position-sticky top-0 border-end d-flex flex-column">
      <Files currentFile={currentFile} setCurrentFile={setCurrentFile} />
      <ToC toc={toc} />
    </aside>
  );
};

export default Sidebar;