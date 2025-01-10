import React from 'react';

const GridContainer = ({ children, cols, customGap, inline, customPadding }) => {
    const gap = customGap !== undefined ? `gap-${customGap}` : 'gap-7';
    const padding = customPadding !== undefined ? `p-${customPadding}` : 'p-2';
    const layoutClass = inline ? 'inline-flex w-full' : `grid grid-cols-${cols}`;

    return (
        <div className={`container ${padding}`}>
            <section className={`${layoutClass} ${gap} text-center`}>
                {children}
            </section>
        </div>
    );
};

export default GridContainer;
