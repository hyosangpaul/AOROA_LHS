-- 요구사항:
-- 1. ETD 기간: 2025년 5월 1일 ~ 2025년 5월 11일
-- 2. POL(출발항) + 컨테이너 타입(CNTR_TYPE) 별로 집계
-- 3. 컨테이너 수량(CNTR 개수) 및 총 중량 합계(CNTR_WGT) 조회
-- 4. 결과 컬럼: POL_CD, CNTR_TYPE, CNTR_COUNT, TOTAL_WGT
-- 5. 정렬: POL_CD, CNTR_TYPE

-- 여기에 SQL 쿼리를 작성하세요

SELECT H.POL_CD, C.CNTR_TYPE, 
    COUNT(C.CNTR_NO) AS CNTR_COUNT, 
    SUM(C.CNTR_WGT) AS TOTAL_WGT
FROM FMS_HBL_MST H
    JOIN FMS_HBL_CNTR C 
        ON H.HBL_NO = C.HBL_NO
            WHERE H.ETD BETWEEN '20250501' AND '20250511'
                GROUP BY H.POL_CD, C.CNTR_TYPE
                ORDER BY H.POL_CD, C.CNTR_TYPE;