import { AutoLayout } from "../../components/autolayout";
import styles from "./styles.module.scss"

import frameSend from "../../assets/frameSend.svg"

const Money = () => {
    return (
        <AutoLayout>
            <div className={styles.money}>
                <div className={styles.money__cards}>
                    <h2>Карты</h2>
                    <div className={styles.money__cardsRow}>
                        <div className={styles.money__cards__card}>
                            <div className={styles.money__cards__card__rectangle}></div>
                            <div className={styles.money__cards__card__content}>
                                <div className={styles.money__cards__card__content__title}>
                                    <span>Привилегия</span>
                                    <span>**** 4319</span>
                                </div>
                                <div className={styles.money__cards__card__content__info}>
                                    <span>54722</span>
                                    <span>₽</span>
                                </div>
                                <div className={styles.money__cards__card__content__ad}>
                                    <span>кешбэк 2% на все</span>
                                    <span>до 31 окт</span>
                                </div>
                            </div>
                        </div>
                        <div className={styles.money__cards__card}>
                            <div className={styles.money__cards__card__rectangle}></div>
                            <div className={styles.money__cards__card__content}>
                                <div className={styles.money__cards__card__content__title}>
                                    <span>Дебетовая</span>
                                    <span>**** 5637</span>
                                </div>
                                <div className={styles.money__cards__card__content__info}>
                                    <span>32100</span>
                                    <span>₽</span>
                                </div>
                                <div className={styles.money__cards__card__content__ad}>
                                    <span>кешбэк 1% на все</span>
                                    <span>до 31 янв</span>
                                </div>
                            </div>
                        </div>
                        <div className={styles.money__cards__add}>
                            <div className={styles.money__cards__add__frame}>
                                +
                            </div>
                        </div>
                    </div>
                </div>
                <div className={styles.money__cards}>
                    <h2>Кредиты и вклады</h2>
                    <div className={styles.money__cardsRow}>
                        <div className={styles.money__cards__card}>
                            <div className={styles.money__cards__card__rectangle}></div>
                            <div className={styles.money__cards__card__content}>
                                <div className={styles.money__cards__card__content__title}>
                                    <span>Вклад</span>
                                    <span>11%</span>
                                </div>
                                <div className={styles.money__cards__card__content__info}>
                                    <span>54722</span>
                                    <span>₽</span>
                                </div>
                                <div className={styles.money__cards__card__content__ad}>
                                    <span>выплатим 5062₽</span>
                                    <span>до 31 окт</span>
                                </div>
                            </div>
                        </div>
                        <div className={styles.money__cards__card}>
                            <div className={styles.money__cards__card__rectangle}></div>
                            <div className={styles.money__cards__card__content}>
                                <div className={styles.money__cards__card__content__title}>
                                    <span>Кредит</span>
                                    <span>12%</span>
                                </div>
                                <div className={styles.money__cards__card__content__info}>
                                    <span>32100</span>
                                    <span>₽</span>
                                </div>
                                <div className={styles.money__cards__card__content__ad}>
                                    <span>осталось 3 мес</span>
                                    <span>до 31 янв</span>
                                </div>
                            </div>
                        </div>
                        <div className={styles.money__cards__add}>
                            <div className={styles.money__cards__add__frame}>
                                +
                            </div>
                        </div>
                    </div>
                </div>
                <div className={styles.money__action}>
                    <h2>Действия</h2>
                    <button className={styles.money__action__send}>
                        <p>Перевести</p>
                        <img src={frameSend} alt="" width={"36px"} height={"36px"} />
                    </button>
                </div>
                <div className={styles.money__payments}>
                    <h2>Активные платежи</h2>
                    <div className={styles.money__payments__cards}>
                        <div className={styles.money__payments__cards__card}>
                            <div className={styles.money__payments__cards__card__rectangle}></div>
                            <div className={styles.money__payments__cards__card__content}>
                                <div className={styles.money__payments__cards__card__content__title}>
                                    <span>Сервисы Яндекса</span>
                                    <span>Оплата по Яндекс Сплит</span>
                                </div>
                                <div className={styles.money__payments__cards__card__content__info}>
                                    <span>6024</span>
                                    <span>₽</span>
                                </div>
                                <div className={styles.money__payments__cards__card__content__ad}>
                                    <div className={styles.money__payments__cards__card__content__ad__text}>
                                        <span>осталось 3012</span>
                                        <span>до 31 окт</span>
                                    </div>
                                    <button>Оплатить</button>
                                </div>
                            </div>
                        </div>
                        <div className={styles.money__payments__cards__card}>
                            <div className={styles.money__payments__cards__card__rectangle}></div>
                            <div className={styles.money__payments__cards__card__content}>
                                <div className={styles.money__payments__cards__card__content__title}>
                                    <span>ГИБДД</span>
                                    <span>Штраф</span>
                                </div>
                                <div className={styles.money__payments__cards__card__content__info}>
                                    <span>1500</span>
                                    <span>₽</span>
                                </div>
                                <div className={styles.money__payments__cards__card__content__ad}>
                                    <div className={styles.money__payments__cards__card__content__ad__text}>
                                        <span>можно оплатить</span>
                                        <span>до 31 окт</span>
                                    </div>
                                    <button>Оплатить</button>
                                </div>
                            </div>
                        </div>
                        <div className={styles.money__payments__cards__card}>
                            <div className={styles.money__payments__cards__card__rectangle}></div>
                            <div className={styles.money__payments__cards__card__content}>
                                <div className={styles.money__payments__cards__card__content__title}>
                                    <span>t2</span>
                                    <span>Оплата сотовой свзяи</span>
                                </div>
                                <div className={styles.money__payments__cards__card__content__info}>
                                    <span>450</span>
                                    <span>₽</span>
                                </div>
                                <div className={styles.money__payments__cards__card__content__ad}>
                                    <div className={styles.money__payments__cards__card__content__ad__text}>
                                        <span>долг</span>
                                        <span></span>
                                    </div>
                                    <button>Оплатить</button>
                                </div>
                            </div>
                        </div>
                        <div className={styles.money__payments__cards__card}>
                            <div className={styles.money__payments__cards__card__rectangle}></div>
                            <div className={styles.money__payments__cards__card__content}>
                                <div className={styles.money__payments__cards__card__content__title}>
                                    <span>ГИБДД</span>
                                    <span>Штраф</span>
                                </div>
                                <div className={styles.money__payments__cards__card__content__info}>
                                    <span>1500</span>
                                    <span>₽</span>
                                </div>
                                <div className={styles.money__payments__cards__card__content__ad}>
                                    <div className={styles.money__payments__cards__card__content__ad__text}>
                                        <span>можно оплатить</span>
                                        <span>до 31 окт</span>
                                    </div>
                                    <button>Оплатить</button>
                                </div>
                            </div>
                        </div>
                    </div>

                </div>
                <div className={styles.money__cards}>
                    <h2>Кешбэк и бонусы</h2>
                    <div className={styles.money__cardsRow}>
                        <div className={styles.money__cards__card}>
                            <div className={styles.money__cards__card__rectangle}></div>
                            <div className={styles.money__cards__card__content}>
                                <div className={styles.money__cards__card__content__title}>
                                    <span>Бонусы Reliance</span>
                                </div>
                                <div className={styles.money__cards__card__content__info}>
                                    <span>312</span>
                                    <span>бонусов</span>
                                </div>
                                <div className={styles.money__cards__card__content__ad}>
                                    <span></span>
                                    <span>1 бонус = 1₽</span>
                                </div>
                            </div>
                        </div>
                        <div className={styles.money__cards__card}>
                            <div className={styles.money__cards__card__rectangle}></div>
                            <div className={styles.money__cards__card__content}>
                                <div className={styles.money__cards__card__content__title}>
                                    <span>На одежду</span>
                                    <span>1%</span>
                                </div>
                                <div className={styles.money__cards__card__content__info}>
                                    <span>450</span>
                                    <span>₽</span>
                                </div>
                                <div className={styles.money__cards__card__content__ad}>
                                    <span>выплатим</span>
                                    <span>до 31 янв</span>
                                </div>
                            </div>
                        </div>
                        <div className={styles.money__cards__card}>
                            <div className={styles.money__cards__card__rectangle}></div>
                            <div className={styles.money__cards__card__content}>
                                <div className={styles.money__cards__card__content__title}>
                                    <span>Рестораны</span>
                                    <span>1%</span>
                                </div>
                                <div className={styles.money__cards__card__content__info}>
                                    <span>340</span>
                                    <span>₽</span>
                                </div>
                                <div className={styles.money__cards__card__content__ad}>
                                    <span>выплатим</span>
                                    <span>до 31 янв</span>
                                </div>
                            </div>
                        </div>
                        <div className={styles.money__cards__card}>
                            <div className={styles.money__cards__card__rectangle}></div>
                            <div className={styles.money__cards__card__content}>
                                <div className={styles.money__cards__card__content__title}>
                                    <span>Косметика</span>
                                    <span>15%</span>
                                </div>
                                <div className={styles.money__cards__card__content__info}>
                                    <span>1032</span>
                                    <span>₽</span>
                                </div>
                                <div className={styles.money__cards__card__content__ad}>
                                    <span>выплатим</span>
                                    <span>до 31 янв</span>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <div className={styles.money__history}>
                    <h2>История</h2>
                    <button className={styles.money__action__send}>
                        <p>Перевести</p>
                        <img src={frameSend} alt="" width={"36px"} height={"36px"} />
                    </button>
                </div>
            </div>
        </AutoLayout>

    );
}

export { Money };