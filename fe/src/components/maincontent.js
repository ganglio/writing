import Sidebar from './sidebar';

const MainContent = () => {
  return (
    <main className="container-fluid">
      <div className="row">
        <div className="col-2">
            <Sidebar />
        </div>
        <div className="col-10">
            <h1>Main Content</h1>
            <p>Welcome to the main content area of the Writing application.</p>
        </div>
      </div>
    </main>
  );
};

export default MainContent;