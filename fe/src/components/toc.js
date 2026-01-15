function buildTree(items) {
    const root = [];
    const stack = [{ level: 0, children: root }];

    for (const item of items) {
        const node = { title: item.title, level: item.level, children: [] };

        // Move up until we find the parent level
        while (stack.length > 0 && stack[stack.length - 1].level >= item.level) {
            stack.pop();
        }

        // Parent is now at top of stack
        stack[stack.length - 1].children.push(node);

        // Push this node as the next potential parent
        stack.push(node);
    }

    return root;
}

const renderTree = (nodes) => {
  return (
    <ul className="list-group list-group-flush">
      {nodes.map((node, index) => (
        <li key={index} className="list-group-item">
          <a className="nav-link" href={`#${node.title}`}>{node.title}</a>
          {node.children.length > 0 && renderTree(node.children)}
        </li>
      ))}
    </ul>
  );
}

const ToC = ({ toc }) => {
  console.log(toc);
  const tree = buildTree(toc);

  return (
    <aside className="fs-6">
      <span className="fw-bold">Table of Contents</span>
      {renderTree(tree)}
    </aside>
  );
};

export default ToC;