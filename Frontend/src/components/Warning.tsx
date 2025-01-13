import { Link } from "react-router-dom";

interface IWarningProps {
    label: string;
    linkText: string,
    to: string
}

export function Warning({label, linkText, to}: IWarningProps){
    return<div className="py-2 text-sm flex justify-center text-slate-400">
        <div>
            {label}
        </div>
        <Link to={to} className="underline pl-1 hover:text-blue-500 hover:font-semibold">
            {linkText}
        </Link>
    </div>
}