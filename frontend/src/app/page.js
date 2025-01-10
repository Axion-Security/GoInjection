"use client";
import {Link, Divider} from "@nextui-org/react";
import {Icon} from "@iconify/react";
import {BrowserOpenURL} from "../../wailsjs/runtime";


export default function Page() {
  return (
      <div className={`flex flex-col justify-center items-center p-12`}>
          <p className={`text-2xl font-bold my-2`}>Welcome to GoInjection!</p>
          <div className={`flex flex-col items-center justify-center`}>
              <div className={`inline-flex`}>
                  <Icon icon="solar:double-alt-arrow-left-outline"/>
                  <Divider className={`my-2 w-[600px]`}/>
                  <Icon icon="solar:double-alt-arrow-right-outline"/>
              </div>
              <p className={`text-sm text-zinc-500`}>Developed By <b className={`bg-gradient-to-r from-blue-500 to-purple-500 bg-clip-text text-transparent underline`}>Fourier</b> for <Link
                  className={`cursor-default`} onPress={() => BrowserOpenURL("https://github.com/Axion-Security")}>Axion
                  Security</Link> </p>
          </div>
      </div>
  )
}