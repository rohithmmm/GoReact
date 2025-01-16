import { Icon } from "../assets/icon";

export function Header() {
    return (
        <header className="fixed top-0 right-0 left-0 border-b border-gray-800 p-4 bg-black">
            <div className="text-5xl text-white font-bold flex justify-center items-center">
                <div className="p-2"><Icon /></div>
                <div>Timesheet</div>
            </div>
        </header>
    )
}