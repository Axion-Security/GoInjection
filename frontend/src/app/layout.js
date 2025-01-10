"use client";
import { Geist, Geist_Mono } from "next/font/google";
import "./globals.css";
import Link from "next/link";
import { Icon } from "@iconify/react";
import {BreadcrumbItem, Breadcrumbs, Button, Divider, Input, Tooltip} from "@nextui-org/react";
import {useEffect, useState} from "react";
import {usePathname, useRouter} from "next/navigation";
import {Quit, WindowMinimise} from "../../wailsjs/runtime";
import Blocker from "../components/ContextMenuHandler";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

const linksList = [
    { link: "/", label: "Dashboard", icon: "solar:home-2-bold-duotone", bottom: false },
    { link: "/waf", label: "WAF Identifier", icon: "solar:shield-keyhole-bold-duotone", bottom: false },
    { link: "/fingerprint", label: "Fingerprint", icon: "solar:database-bold-duotone", bottom: false },
    { link: "/interpreter", label: "Interpreter", icon: "solar:code-circle-bold-duotone", bottom: false },
    { link: "/resolver", label: "Resolver", icon: "solar:magnifer-bold-duotone", bottom: false },
    { link: "/injector", label: "Injector", icon: "solar:target-bold-duotone", bottom: false },
    { link: "/infos", label: "Target Information's", icon: "solar:info-circle-bold-duotone", bottom: true },
];

const programName = "GoInjection";
const programIcon = "solar:syringe-bold-duotone";

export default function RootLayout({ children }) {
    const asPath = usePathname();
    const [breadcrumbs, setBreadcrumbs] = useState([]);

    useEffect(() => {
        const matchedLink = linksList.find((link) => link.link === asPath);
        if (matchedLink) {
            setBreadcrumbs([{ label: programName }, { label: matchedLink.label }]);
        } else {
            setBreadcrumbs([{ label: programName }, { label: "Not Found :(" }]);
        }
    }, [asPath]);

    //Blocker()

  return (
    <html lang="en" className={`dark h-full overflow-hidden select-none`}>
      <body
        className={`${geistSans.variable} ${geistMono.variable} antialiased flex flex-col h-screen`}
      >
      <Navbar />
      <div className="flex flex-1">
          <div className={`p-4`}>
              <Sidebar />
          </div>

          <main className="flex-1 pr-2 overflow-y-auto mt-4">
              <div className="bg-[#2c2d31] text-white p-2 draggable rounded mb-1">
                  <Breadcrumbs isDisabled>
                      {breadcrumbs.map((breadcrumb, index) => (
                          <BreadcrumbItem key={index}>{breadcrumb.label}</BreadcrumbItem>
                      ))}
                  </Breadcrumbs>
              </div>
              <div className={`bg-[#2c2d31] rounded p-2`}>
                  {children}
              </div>
          </main>
      </div>
      </body>
    </html>
  );
}

function Sidebar() {
    return (
        <aside className="flex flex-col justify-between bg-[#2c2d31] text-white p-4 h-full rounded-lg draggable">
            <div>
            <div className="w-12 h-12 overflow-hidden rounded-full bg-gradient-to-r from-blue-500 to-blue-800 relative">
                    <Icon icon={programIcon} className="w-full h-full text-[#e7e7e7] -scale-75 rotate-180" alt="Logo" />
                </div>
                <Divider className="my-4" />

                <div>
                    {linksList.map((button, index) => (
                        !button.bottom && (
                            <AnimatedLink key={index} iconName={button.icon} link={button.link} label={button.label} />
                        )
                    ))}

                </div>
            </div>

            <div className={`text-center`}>
                {linksList.map((button, index) => (
                    button.bottom && (
                        <AnimatedLink key={index} iconName={button.icon} link={button.link} label={button.label} />
                    )
                ))}

            </div>
        </aside>
    )
}

function Navbar() {
    return (
        <header className="bg-[#2c2d31] text-white p-2 draggable">
            <nav className="mx-3 flex items-center justify-between">
                <div className={``}>
                    <Icon className={`h-auto mr-3`} height={25} icon={"solar:programming-bold-duotone"} />
                </div>

                <div className={"inline-flex"}>
                    <p>{programName}</p>
                </div>

                <div className="flex items-center">
                    <button
                        onClick={() => WindowMinimise()}
                        className="w-4 h-4 bg-yellow-500 hover:bg-yellow-600 rounded-full transition-all duration-300 mr-2"
                    />

                    <button
                        onClick={() => Quit()}
                        className="w-4 h-4 bg-red-500 hover:bg-red-600 rounded-full transition-all duration-300"
                    />
                </div>
            </nav>
        </header>

    )
}

function AnimatedLink({ link, iconName, label }) {
    const asPath = usePathname();
    const isActive = link === asPath;

    return (
        <Tooltip closeDelay={200} showArrow={true} placement={"right"} content={label}>
            <Link
                prefetch={true}
                href={link}
                className="w-full outline-0"
            >
                <div
                    className={`mt-1 w-full flex items-center justify-center rounded-lg ${
                        isActive && "bg-gradient-to-br from-blue-500 to-[#374969]"
                    } hover:bg-[#3e4046] hover:text-white hover:scale-105 p-2 cursor-pointer transition-all duration-300`}
                >
                    <Icon icon={iconName} className="w-8 h-8" />
                </div>
            </Link>
        </Tooltip>
    );
}