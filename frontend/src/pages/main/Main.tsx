import React from "react";
import styles from "./styles.module.scss";
import classNames from "classnames";

import man from "../../assets/young_man.png";
import frameInv from "../../assets/frameInv.svg";
import frameCredit from "../../assets/frameCredit.svg";
import frameDeposit from "../../assets/frameDeposit.svg";
import framePremium from "../../assets/framePremium.svg";
import debet from "../../assets/debet.svg"
import vip from "../../assets/vip.svg"
import sticker from "../../assets/sticker.svg"
import sticker_m from "../../assets/sticker_m.svg"
import card_m from "../../assets/card_m.svg"
import for_close from "../../assets/for_close.svg"
import junior from "../../assets/junior.svg"
import service1 from "../../assets/service1.svg"
import service2 from "../../assets/service2.svg"
import service3 from "../../assets/service3.svg"
import { AutoLayout } from "../../components/autolayout";


const Main = () => {
    return (
        <AutoLayout>
            <div className={styles.main} id="main">
                <h2>Мы - Reliance</h2>
                <div className={styles.main__banner}>
                    <div className={styles.main__banner__info}>
                        <h3>Добро пожаловать</h3>
                        <p>Хорошего дня!</p>
                        <button>Войти</button>
                    </div>
                    <svg xmlns="http://www.w3.org/2000/svg" width="520" height="302" viewBox="0 0 520 302" fill="none">
                        <path opacity="0.1" d="M1 300.965V1.41375M1.71982 1H520M1.71982 20.9698H520M1.71982 40.9395H520M1.71982 60.9093H520M1.71982 80.8791H520M1.71982 100.849H520M1.71982 120.819H520M1.71982 140.788H520M1.71982 160.758H520M1.71982 180.728H520M1.71982 200.698H520M1.71982 220.668H520M1.71982 240.637H520M1.71982 260.607H520M1.71982 280.577H520M1.71982 300.547H520M35.5515 300.967V1.41609M70.103 300.97V1.41842M104.654 300.972V1.42076M139.206 300.974V1.4231M173.757 300.977V1.42543M208.309 300.979V1.42777M242.86 300.981V1.4301M277.412 300.984V1.43244M311.963 300.986V1.43478M346.515 300.988V1.43711M381.066 300.991V1.43945M415.618 300.993V1.44179M450.169 300.995V1.44412M484.721 300.998V1.44646M519.272 301V1.44879" stroke="url(#paint0_linear_80_77)" />
                        <defs>
                            <linearGradient id="paint0_linear_80_77" x1="-21.7542" y1="165.586" x2="564.175" y2="174.539" gradientUnits="userSpaceOnUse">
                                <stop offset="0.24" stop-color="#0D6AEF" />
                                <stop offset="0.48" stop-color="white" />
                                <stop offset="0.715" stop-color="#307CEA" />
                            </linearGradient>
                        </defs>
                    </svg>
                    <div className={styles.main__banner__img}>
                        <img src={man} alt="" />
                    </div>
                </div>
                <div className={styles.main__info}>
                    <div className={styles.main__info__card}>
                        <div className={classNames(styles.main__info__card__rectangle, styles.main__info__card__investRectangle)}></div>
                        <div className={styles.main__info__card__description}>
                            <p className={styles.main__info__card__name}>Инвестиции</p>
                            <span className={classNames(styles.main__info__card__text, styles.main__info__card__invest)}>543 </span>
                            <span className={classNames(styles.main__info__card__textSmall, styles.main__info__card__invest)}>акций</span>
                            <p className={styles.main__info__card__stat}>
                                <svg xmlns="http://www.w3.org/2000/svg" width="12" height="12" viewBox="0 0 12 12" fill="none">
                                    <rect width="12" height="12" rx="2" fill="#0B68EF" fill-opacity="0.15" />
                                    <path d="M10 4L6.6 7.4L4.6 5.4L2 8M10 4H7.6M10 4V6.4" stroke="#0B68EF" stroke-linecap="round" stroke-linejoin="round" />
                                </svg>
                                <span className={classNames(styles.main__info__card__statNumber, styles.main__info__card__invest)}>17%</span>
                                <span className={styles.main__info__card__statText}>рост за год</span>
                            </p>
                        </div>
                        <img src={frameInv} alt="" />
                    </div>
                    <div className={styles.main__info__card}>
                        <div className={classNames(styles.main__info__card__rectangle, styles.main__info__card__creditRectangle)}></div>
                        <div className={styles.main__info__card__description}>
                            <p className={styles.main__info__card__name}>Кредиты</p>
                            <span className={classNames(styles.main__info__card__textSmall, styles.main__info__card__credit)}>от</span>
                            <span className={classNames(styles.main__info__card__text, styles.main__info__card__credit)}> 20%</span>
                            <p className={styles.main__info__card__stat}>
                                <svg xmlns="http://www.w3.org/2000/svg" width="12" height="12" viewBox="0 0 12 12" fill="none">
                                    <rect width="12" height="12" rx="2" fill="#FFC737" fill-opacity="0.15" />
                                    <path d="M10 4L6.6 7.4L4.6 5.4L2 8M10 4H7.6M10 4V6.4" stroke="#FFC737" stroke-linecap="round" stroke-linejoin="round" />
                                </svg>
                                <span className={classNames(styles.main__info__card__statNumber, styles.main__info__card__credit)}>112</span>
                                <span className={styles.main__info__card__statText}>выдано за неделю</span>
                            </p>
                        </div>
                        <img src={frameCredit} alt="" />
                    </div>
                    <div className={styles.main__info__card}>
                        <div className={classNames(styles.main__info__card__rectangle, styles.main__info__card__depositRectangle)}></div>
                        <div className={styles.main__info__card__description}>
                            <p className={styles.main__info__card__name}>Вклады</p>
                            <span className={classNames(styles.main__info__card__textSmall, styles.main__info__card__deposit)}>до</span>
                            <span className={classNames(styles.main__info__card__text, styles.main__info__card__deposit)}> 21%</span>
                            <p className={styles.main__info__card__stat}>
                                <svg xmlns="http://www.w3.org/2000/svg" width="12" height="12" viewBox="0 0 12 12" fill="none">
                                    <rect width="12" height="12" rx="2" fill="#2F887C" fill-opacity="0.15" />
                                    <path d="M10 4L6.6 7.4L4.6 5.4L2 8M10 4H7.6M10 4V6.4" stroke="#2F887C" stroke-linecap="round" stroke-linejoin="round" />
                                </svg>
                                <span className={classNames(styles.main__info__card__statNumber, styles.main__info__card__deposit)}>768</span>
                                <span className={styles.main__info__card__statText}>открыто за месяц</span>
                            </p>
                        </div>
                        <img src={frameDeposit} alt="" />
                    </div>
                    <div className={styles.main__info__card}>
                        <div className={classNames(styles.main__info__card__rectangle, styles.main__info__card__premiumRectangle)}></div>
                        <div className={styles.main__info__card__description}>
                            <p className={styles.main__info__card__name}>Привилегия</p>
                            <span className={classNames(styles.main__info__card__text, styles.main__info__card__premium)}>1.045 </span>
                            <span className={classNames(styles.main__info__card__textSmall, styles.main__info__card__textSmall__premiunm, styles.main__info__card__premium)}>клиентов</span>
                            <p className={styles.main__info__card__stat}>
                                <svg xmlns="http://www.w3.org/2000/svg" width="12" height="12" viewBox="0 0 12 12" fill="none">
                                    <rect width="12" height="12" rx="2" fill="#FF4E79" fill-opacity="0.15" />
                                    <path d="M10 4L6.6 7.4L4.6 5.4L2 8M10 4H7.6M10 4V6.4" stroke="#FF4E79" stroke-linecap="round" stroke-linejoin="round" />
                                </svg>
                                <span className={classNames(styles.main__info__card__statNumber, styles.main__info__card__premium)}>13</span>
                                <span className={styles.main__info__card__statText}>клиентов в день</span>
                            </p>
                        </div>
                        <img src={framePremium} alt="" />
                    </div>
                </div>

                <div className={styles.main__products} id="products">
                    <h2>Продукты</h2>
                    <div className={styles.main__products__cards}>
                        <div className={styles.main__products__cards__row}>
                            <img src={debet} alt="" />
                            <img src={vip} alt="" />
                            <img src={card_m} alt="" />
                        </div>

                        <div className={styles.main__products__cards__row}>
                            <img src={sticker} alt="" />
                            <img src={sticker_m} alt="" />
                            <img src={for_close} alt="" />
                            <img src={junior} alt="" />
                        </div>
                    </div>
                    <div className={styles.main__products__button}>
                        <button>Заказать</button>
                    </div>

                </div>

                <div className={styles.main__services} id="services">
                    <h2>Сервисы и услуги</h2>
                    <div className={styles.main__services__cards}>
                        <img src={service1} alt="" />
                        <img src={service2} alt="" />
                        <img src={service3} alt="" />
                    </div>
                </div>
            </div>
        </AutoLayout>
    );
}

export { Main };