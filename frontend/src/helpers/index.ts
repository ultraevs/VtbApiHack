import Cookies from "js-cookie";

export const convertFileToBase64 = (file: File): Promise<string> => {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();

    reader.onload = () => {
      const base64 = reader.result as string;
      resolve(base64);
    };

    reader.onerror = (error) => reject(error);

    reader.readAsDataURL(file);
  });
};

export const setAuthTokenCookie = (token: string) => {
  Cookies.set("Authtoken", token, { expires: 7 });
};