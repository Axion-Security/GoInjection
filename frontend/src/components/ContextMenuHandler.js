import { useEffect } from 'react';

export default function Blocker() {
    useEffect(() => {
        const disableRightClick = (event) => {
            event.preventDefault();
        };
        document.addEventListener('contextmenu', disableRightClick);
        return () => {
            document.removeEventListener('contextmenu', disableRightClick);
        };
    }, []);
};
