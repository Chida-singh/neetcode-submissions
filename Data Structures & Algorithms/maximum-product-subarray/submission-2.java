class Solution {
    public int maxProduct(int[] nums) {
        int maxProd = nums[0];
        int suffix = 1;
        int prefix = 1;

        int n = nums.length;

        for(int i =0; i< n;i++){
            prefix*= nums[i];
            suffix*= nums[n-i-1];

            maxProd = Math.max(maxProd, Math.max(suffix, prefix));

            if(nums[i] == 0) prefix = 1;
            if(nums[n-i-1] == 0) suffix = 1;
        }
        return maxProd;
    }
}
