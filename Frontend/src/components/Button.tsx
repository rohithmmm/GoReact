
interface IButtonProps {
    label: string;
    onClick: () => void;
    disabled?: boolean;
}

export function Button({label, onClick, disabled}: IButtonProps){
    return<button disabled={disabled} className="w-full px-5 py-3 mb-2 font-medium bg-blue-600 text-white rounded-lg disabled:bg-blue-500" onClick={onClick}>
        {label}
    </button>
}