import axios from "axios";
import { RegisterInput } from "../../../consts/types";

export const createUser = async (input: RegisterInput) => {
  try {
    const response = await axios.post(
      "https://tatarby.shmyaks.ru/v1/user_create",
      {
        surname: input.surname,
        email: input.email,
        password: input.password,
      },
      {
        headers: {
          "Content-Type": "application/json",
        },
      }
    );
    return { success: true, data: response.data.token };
  } catch (error) {
    return {success: false, error: error};
  }
};