-- 요구사항:
-- 1. HOUSE B/L (HBL_NO) 를 기준으로 각 B/L 별 컨테이너 수량(CNTR 개수) 을 집계
-- 2. 가장 많은 수량을 가진 B/L 1건을 조회
-- 3. 컨테이너 수량은 CNTR_NO 기준으로 COUNT
-- 4. 동일 수량 시 ETD 빠른 순으로 우선 선택
-- 5. 결과 컬럼: HBL_NO, CNTR_COUNT, ETD
-- 6. 정렬: CNTR_COUNT DESC, ETD ASC

-- 여기에 SQL 쿼리를 작성하세요

SELECT HBL_NO, CNTR_COUNT, ETD
FROM (
    SELECT H.FBL_NO, 
        COUNT(C.CNTR_NO) AS CNTR_COUNT, H.ETD,
        RANK() OVER (
            ORDER BY COUNT(C.CNTR_NO) DESC, H.ETD ASC) AS CNTR_COUNT_RANK
    FROM FMS_HBL_MST H
        JOIN FMS_HBL_CNTR C 
            ON H.HBL_NO = C.HBL_NO
                GROUP BY H.HBL_NO, H.ETD
)
WHERE CNTR_COUNT_RANK = 1;