import { mainFn2 } from './asyncTasts';
/**
 * Adds two numbers.
 * @param {number} x - The first number.
 * @param {number} y - The second number.
 * @returns {number} The sum of the two numbers.
 */
export const mainFn = (x, y) =>{
    return x + y;
  }

  /**
   * @param {Omit<import('./asyncTasts').Profile, 'age'>} main
   * @param {Order} order
   */
  export const test = (main, order)=>{
    // '5'.padStart(2, '0'); // returns '05'
  }