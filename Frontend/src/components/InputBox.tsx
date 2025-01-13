import { RefObject } from "react"

interface IInputBox {
    label: string, 
    placeholder: string,
    inputRef?: RefObject<HTMLInputElement>
}
export function InputBox({label, placeholder, inputRef}: IInputBox){
    return<div>
        <div className="py-2 text-sm font-medium text-left text-white">
            {label}
        </div>
        <input ref={inputRef} placeholder={placeholder} className="w-full px-2 py-2 border rounded-md font-semibold"/>
    </div>
}