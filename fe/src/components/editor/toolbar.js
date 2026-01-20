const toolbarButtons = [
    [
        { icon: 'bi-type-bold', alt: 'bold'},
        { icon: 'bi-type-italic', alt: 'italic' },
        { icon: 'bi-type-underline', alt: 'underline' },      
    ],
    [
        { icon: 'bi-type-h1', alt: 'heading1' },
        { icon: 'bi-type-h2', alt: 'heading2' },
        { icon: 'bi-type-h3', alt: 'heading3' },
        { icon: 'bi-type-h4', alt: 'heading4' },
        { icon: 'bi-type-h5', alt: 'heading5' },
        { icon: 'bi-type-h6', alt: 'heading6' },
    ]
    // [
    //     { icon: 'bi-link-45deg', alt: 'link' },
    //     { icon: 'bi-image', alt: 'image' },
    // ]  
];

const actions = {
    bold: (text) => `**${text}**`,
    italic: (text) => `*${text}*`,
    underline: (text) => `__${text}__`,
    heading1: (text) => `# ${text}`,
    heading2: (text) => `## ${text}`,
    heading3: (text) => `### ${text}`,
    heading4: (text) => `#### ${text}`,
    heading5: (text) => `##### ${text}`,
    heading6: (text) => `###### ${text}`,
    // link: (text) => `[${text}](url)`,
    // image: (text) => `![${text}](image-url)`,
};

const Toolbar = ({ selection, content, setContent, isDirty = false }) => {
    return (
        <div className="d-inline-flex justify-content-between w-100">
            <div className="btn-toolbar" role="toolbar" aria-label="Toolbar with button groups">
                {
                    toolbarButtons.map((group, index) => (
                        <div className="btn-group me-2" key={index}>
                            {
                                group.map((button, btnIndex) => (
                                    <button type="button" className="btn btn-sm btn-outline-secondary" key={btnIndex}>
                                        <i className={`bi ${button.icon}`} title={button.alt} onClick={() => {
                                            if (!selection || selection.isCollapsed || selection.type !== 'Range') {
                                                return;
                                            }
                                            const updatedText = actions[button.alt](selection.toString());
                                            const newContent = content.slice(0, selection.anchorOffset) + updatedText + content.slice(selection.focusOffset);
                                            setContent(newContent);
                                        }}></i>
                                    </button>
                                ))
                            }
                        </div>
                    ))
                }
            </div>
            <button className="btn btn-sm btn-primary" disabled={!isDirty}>Save</button>
        </div>
    );
};

export default Toolbar;