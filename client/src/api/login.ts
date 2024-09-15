
import axios from 'axios';
import { hostUrl } from '../utils/util';
import { AxiosResponse } from 'axios';
import { AdminPasswordRes, IsAdminAvailableRes, LoginRes } from '../models/login';



const URL = hostUrl()

export const isAdminAvailable = async():Promise<AxiosResponse<IsAdminAvailableRes>>=>{
    const url = URL+"/api/auth/isAdmin"
    const res = await axios.get<IsAdminAvailableRes>(url)
    return res
}

export const registerAdmin= async(username: string, password:string, initPass: string) =>{
    const url = URL+"/api/auth/init"
    const res = await axios.post(url, {
        username: username,
        password: password,
        initPassword: initPass
    })
    return res
}

export const login= async(username: string, password: string):Promise<AxiosResponse<LoginRes>>=>{
    const url = URL+"/api/auth/login"
    const res = await axios.post<LoginRes>(url, {
        username: username, 
        password: password
    })
    return res
}

export const adminPassword= async(initPass: string, password: string):Promise<AxiosResponse<AdminPasswordRes>> =>{
    const url = URL+"/api/auth/getAdminPassword"
    const res = await axios.get<AdminPasswordRes>(url)
    return res
}