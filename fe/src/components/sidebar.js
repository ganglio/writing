import Files from "./files";
import ToC from "./toc";

const Sidebar = () => {
  return (
    <aside className="position-sticky top-0 border-end d-flex flex-column">
      <Files />
      <ToC />
    </aside>
  );
};

export default Sidebar;