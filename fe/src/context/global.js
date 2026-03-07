import { createContext, useReducer } from "react";

const Actions = {
	SET_CURRENT_OPEN_FILE: "SET_CURRENT_OPEN_FILE",
	SET_CURRENT_OPEN_FILE_CONTENT: "SET_CURRENT_OPEN_FILE_CONTENT",
	SET_FILE_LIST: "SET_FILE_LIST",
	SET_TOC: "SET_TOC",
	TOGGLE_MODAL: "TOGGLE_MODAL",
}

const InitalState = {
	currentOpenFile: null,
	currentOpenFileContent: null,
	files: [],
	toc: [],
	modal: null,
}

const Reducer = (state, action) => {
	console.log("Reducer called with action:", action, state);
	switch (action.type) {
		case Actions.SET_CURRENT_OPEN_FILE:
		return {
			...state,
			currentOpenFile: action.payload,
	};
	case Actions.SET_CURRENT_OPEN_FILE_CONTENT:
	  	return {
			...state,
			currentOpenFileContent: action.payload,
	  	};
	case Actions.SET_FILE_LIST:
	  	return {
			...state,
			files: action.payload,
	  	};
	case Actions.SET_TOC:
	  	return {
			...state,
			toc: action.payload,
	  	};
	case Actions.TOGGLE_MODAL:
	  	return {
			...state,
			modal: state.modal ? null : action.payload,
	  	};
	default:
	  	return state;
  }
};

const Context = createContext();

const Provider = ({ children }) => {
	const [state, dispatch] = useReducer(Reducer, InitalState);

	return (
		<Context.Provider value={{ state, dispatch }}>
			{children}
		</Context.Provider>
	);
};

export default {
	Provider,
	Context,
	Actions
};
