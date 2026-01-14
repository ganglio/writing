const Footer = () => {
    const year = new Date().getFullYear();
    return (
        <nav className="navbar bg-body-tertiary">
            <div className="container-fluid">
                <span className="navbar-text mx-auto py-0 fs-6">&copy; {year} Garatsu Inc. All rights reserved.</span>
            </div>
        </nav>
    );
};

export default Footer;